package repository

import (
	"database/sql"
	"fmt"
	"jobqueue/Internals/models"
)

type JobRepo struct {
	DB *sql.DB
}

func (r *JobRepo) CreateJob(job *models.Job) (int, error) {
	fmt.Printf("Trying to insert job: %+v\n", job)
	result, err := r.DB.Exec(`INSERT INTO jobs (payload, status) VALUES (?, ?)`, job.Payload, job.Status)
	if err != nil {
		fmt.Printf("DB Insert Error: %v\n", err)
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func (r *JobRepo) GetJobById(id int) (*models.Job, error) {
	var job models.Job
	err := r.DB.QueryRow(`SELECT id, payload, status, result, created_at, updated_at FROM jobs WHERE id = ?`,
		id).Scan(&job.ID, &job.Payload, &job.Status, &job.Result, &job.CreatedAt, &job.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *JobRepo) GetAllJobs() ([]*models.Job, error) {
	rows, err := r.DB.Query(`SELECT id, payload, status, result, created_at, updated_at FROM jobs ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []*models.Job
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.ID, &job.Payload, &job.Status, &job.Result, &job.CreatedAt, &job.UpdatedAt); err != nil {
			return nil, err
		}
		jobs = append(jobs, &job)
	}
	return jobs, nil
}

func (r *JobRepo) UpdateStatusAndResult(id int, status string, result string) error {
	_, err := r.DB.Exec(`UPDATE jobs SET status = ?, result = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`, status, result, id)
	return err
}
