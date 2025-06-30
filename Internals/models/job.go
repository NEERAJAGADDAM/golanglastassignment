package models

import "time"

type Job struct {
	ID        int       `json:"id"`
	Payload   string    `json:"payload"`
	Status    string    `json:"status"`
	Result    string    `json:"result,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
