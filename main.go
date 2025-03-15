package main

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"deskription"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}
