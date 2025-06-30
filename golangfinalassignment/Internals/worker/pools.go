package worker

import (
	"jobqueue/Internals/repository"
	"jobqueue/Internals/utils"
	"time"
)

type Job struct {
	ID      int
	Payload string
}

var JobQueue = make(chan Job, 100)

type JobExecutor struct {
	Repo repository.DbRepository
}

func (w *JobExecutor) StartWorkerPool(n int) {
	for i := 0; i < n; i++ {
		go w.worker(i)
	}
}

func (w *JobExecutor) worker(id int) {
	for job := range JobQueue {
		utils.Log.Infof("Worker %d started job %d", id, job.ID)

		time.Sleep(3 * time.Second) // simulate work

		result := "Processed: " + job.Payload
		err := w.Repo.UpdateStatusAndResult(job.ID, "completed", result)
		if err != nil {
			utils.Log.Errorf("Worker %d failed job %d: %v", id, job.ID, err)
		} else {
			utils.Log.Infof("Worker %d finished job %d", id, job.ID)
		}
	}
}
