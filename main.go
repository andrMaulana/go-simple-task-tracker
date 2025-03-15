package main

import (
	"encoding/json"
	"os"
	"time"
)

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

// membuat fungsi load file `json`
func loadTasks() (TaskList, error) {
	var taskList TaskList
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return taskList, err
	}
	json.Unmarshal(data, &taskList)
	return taskList, nil
}
