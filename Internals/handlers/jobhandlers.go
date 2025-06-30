package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"jobqueue/Internals/models"
	"jobqueue/Internals/repository"
	"jobqueue/Internals/worker"

	"github.com/gorilla/mux"
)

type JobHandler struct {
	JobRepo  repository.DbRepository
	JobQueue chan worker.Job
}

func (h *JobHandler) SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Payload string `json:"payload"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	job := &models.Job{
		Payload: input.Payload,
		Status:  "queued",
	}

	id, err := h.JobRepo.CreateJob(job)
	if err != nil {
		http.Error(w, "Failed to create job", http.StatusInternalServerError)
		return
	}

	h.JobQueue <- worker.Job{ID: id, Payload: input.Payload}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Job submitted",
		"job_id":  id,
	})
}

func (h *JobHandler) GetJobHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	job, err := h.JobRepo.GetJobById(id)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

func (h *JobHandler) ListJobsHandler(w http.ResponseWriter, r *http.Request) {
	jobs, err := h.JobRepo.GetAllJobs()
	if err != nil {
		http.Error(w, "Failed to list jobs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}
