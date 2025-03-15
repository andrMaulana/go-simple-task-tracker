package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
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

// fungsi untuk menulis ke file json
func saveTask(taskList TaskList) error {
	data, _ := json.MarshalIndent(taskList, "", "  ")
	return os.WriteFile("tasks.json", data, 0o644)
}

// fungsi untuk memastikan file `json` ada
func ensureFileExists() error {
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		return saveTask(TaskList{Tasks: []Task{}})
	}

	return nil
}

// fungsi untuk melakakukan generate id
func generateID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}

	return tasks[len(tasks)-1].ID + 1
}

var (
	ErrTaskNotFound     = errors.New("tugas tidak ditemukan")
	ErrInvalidStatus    = errors.New("status tidak valid")
	ErrEmptyDeskription = errors.New("deskripsi tidak boleh kosong")
)

func main() {
	ensureFileExists()
	taskList, _ := loadTasks()
	fmt.Printf("%+v\n", taskList.Tasks)
}
