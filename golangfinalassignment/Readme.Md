# Golang Job Queue System

A high-performance Job Queue system built in Go with MySQL, structured logging, worker pool concurrency, RESTful API, and React frontend. Supports async job processing, job status tracking, and deployment to Docker/Render.

---

##  Features

- Submit jobs via REST API
- Track job status and results
- Concurrent processing with a worker pool (5+ workers)
- MySQL backend for persistence
- Structured logging with Zap
- Docker & Docker Compose support
- Render-compatible deployment


---

##  Tech Stack

- **Backend**: Go (net/http, gorilla/mux)
- **Frontend**: React.js
- **Database**: MySQL (via Docker or hosted)
- **Logging**: Uber Zap
- **Deployment**: Docker, Render

---

## Project Structure


golang-jobqueue/
├── cmd/api/main.go # Main entrypoint
├── Internals/
│ ├── config/ # DB config
│ ├── handlers/ # HTTP Handlers
│ ├── models/ # Job model
│ ├── repository/ # DB repository interface
│ ├── utils/ # Logger setup
│ └── worker/ # Worker pool logic
├── Dockerfile
├── docker-compose.yml
├── .env
└── README.md


output:=

create job:
{
    "job_id": 3,
    "message": "Job submitted"
}

get all jobs:

[
    {
        "id": 2,
        "payload": "Process this task",
        "status": "completed",
        "result": "Processed: Process this task",
        "created_at": "2025-06-30T09:56:25Z",
        "updated_at": "2025-06-30T09:56:28Z"
    },
    {
        "id": 1,
        "payload": "Process this task",
        "status": "completed",
        "result": "Processed: Process this task",
        "created_at": "2025-06-30T09:07:37Z",
        "updated_at": "2025-06-30T09:07:40Z"
    }
]

get by id:

{
    "id": 1,
    "payload": "Process this task",
    "status": "completed",
    "result": "Processed: Process this task",
    "created_at": "2025-06-30T09:07:37Z",
    "updated_at": "2025-06-30T09:07:40Z"
}
