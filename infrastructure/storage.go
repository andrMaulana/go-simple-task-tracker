package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/andrMaulana/go-simple-task-tracker/internal/domain"
)

type Storage interface {
	LoadTasks() (domain.TaskList, error)
	SaveTasks(domain.TaskList) error
}

type JsonStorage struct{}

func NewJsonStorage() *JsonStorage {
	return &JsonStorage{}
}

func (s *JsonStorage) LoadTasks() (domain.TaskList, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			// jika file tidak ada, buat file kosong
			emptyTaskList := domain.TaskList{Tasks: []domain.Task{}}
			s.SaveTasks(emptyTaskList)
			return emptyTaskList, nil
		}

		return domain.TaskList{}, err
	}

	var taskList domain.TaskList
	err = json.Unmarshal(data, &taskList)
	if err != nil {
		return domain.TaskList{}, fmt.Errorf("format file tasks.json tidak valid: %v", err)
	}

	return taskList, nil
}

func (s *JsonStorage) SaveTasks(taskList domain.TaskList) error {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0o644)
}
