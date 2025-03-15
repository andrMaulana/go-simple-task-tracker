package infrastructure

import "github.com/andrMaulana/go-simple-task-tracker/internal/domain"

type Storage interface {
	LoadTasks() (domain.TaskList, error)
	SaveTasks(domain.TaskList) error
}

type JsonStorage struct{}

func NewJsonStorage() *JsonStorage {
	return &JsonStorage{}
}
