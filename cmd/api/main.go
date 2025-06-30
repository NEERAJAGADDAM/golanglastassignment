package main

import (
	"log"
	"net/http"
	"os"

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

	// Connect to database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Initialize repository and job handler
	jobRepo := &repository.JobRepo{DB: db}
	jobHandler := &handlers.JobHandler{
		JobRepo:  jobRepo,
		JobQueue: worker.JobQueue,
	}

	// Start worker pool with 5 workers
	jobWorker := &worker.JobExecutor{Repo: jobRepo}
	jobWorker.StartWorkerPool(5)

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
