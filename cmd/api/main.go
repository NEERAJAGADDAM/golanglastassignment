package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"jobqueue/Internals/config"
	"jobqueue/Internals/handlers"
	"jobqueue/Internals/repository"
	"jobqueue/Internals/utils"
	"jobqueue/Internals/worker"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println(" No .env file found. Continuing with system environment variables.")
	}

	// Initialize structured logger
	utils.InitLogger()

	// Retry database connection
	var db *sql.DB
	var err error
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "root:@tcp(db:3306)/jobqueue"
	}

	for i := 0; i < 10; i++ {
		db, err = config.ConnectDB()
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("Successfully connected to the database.")
				break
			}
			log.Printf("Attempt %d: Ping failed: %v", i+1, err)
		} else {
			log.Printf("Attempt %d: Failed to open DB: %v", i+1, err)
		}
		log.Println(" Waiting for DB to be ready...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Database connection failed after multiple attempts: %v", err)
	}
	defer db.Close()

	// Initialize repository and job handler
	jobRepo := &repository.JobRepo{DB: db}
	jobHandler := &handlers.JobHandler{
		JobRepo:  jobRepo,
		JobQueue: worker.JobQueue,
	}

	// Start worker pool with 5 workers
	jobWorker := &worker.JobExecutor{Repo: jobRepo}
	jobWorker.StartWorkerPool(5)

	router := gin.Default()
	// Enable CORS from all origins
	router.Use(cors.Default())

	// Setup API routes
	r := mux.NewRouter()
	r.HandleFunc("/jobs", jobHandler.SubmitJobHandler).Methods("POST")
	r.HandleFunc("/jobs/{id:[0-9]+}", jobHandler.GetJobHandler).Methods("GET")
	r.HandleFunc("/jobs", jobHandler.ListJobsHandler).Methods("GET")

	// Get port from env or fallback to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
