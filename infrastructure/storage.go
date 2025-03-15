package infrastructure

import (
	"encoding/json"
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
	var taskList domain.TaskList
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return taskList, err
	}

	json.Unmarshal(data, &taskList)
	return taskList, nil
}
