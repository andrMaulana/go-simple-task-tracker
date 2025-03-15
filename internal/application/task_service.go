package application

import (
	"errors"
	"time"

	"github.com/andrMaulana/go-simple-task-tracker/infrastructure"
	"github.com/andrMaulana/go-simple-task-tracker/internal/domain"
)

var (
	ErrTaskNotFound     = errors.New("tugas tidak ditemukan")
	ErrInvalidStatus    = errors.New("status tidak valid")
	ErrEmptyDescription = errors.New("deksripksi tidak boleh kosong")
)

type TaskService struct {
	storage infrastructure.Storage
}

func NewTaskService(storage infrastructure.Storage) *TaskService {
	return &TaskService{storage: storage}
}

// method add task
func (s *TaskService) AddTask(description string) (domain.Task, error) {
	if description == "" {
		return domain.Task{}, ErrEmptyDescription
	}

	taskList, _ := s.storage.LoadTasks()
	newTask := domain.Task{
		ID:          generateID(taskList.Tasks),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	taskList.Tasks = append(taskList.Tasks, newTask)
	return newTask, nil
}
