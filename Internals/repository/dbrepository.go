package repository

import "jobqueue/Internals/models"

type DbRepository interface {
	CreateJob(job *models.Job) (int, error)
	GetJobById(id int) (*models.Job, error)
	GetAllJobs() ([]*models.Job, error)
	UpdateStatusAndResult(id int, status string, result string) error
}
