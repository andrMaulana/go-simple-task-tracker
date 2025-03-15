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

// method update task
func (s *TaskService) UpdateTask(id int, description string) error {
	taskList, _ := s.storage.LoadTasks()
	for i, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks[i].Description = description
			taskList.Tasks[i].UpdatedAt = time.Now().UTC()
			s.storage.SaveTasks(taskList)
			return nil
		}
	}

	return ErrTaskNotFound
}

// method update status task
func (s *TaskService) UpdateTaskStatus(id int, status string) error {
	validStatus := map[string]bool{
		"todo":       true,
		"in-progres": true,
		"done":       true,
	}

	if !validStatus[status] {
		return ErrInvalidStatus
	}

	taskList, _ := s.storage.LoadTasks()
	for i, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks[i].Status = status
			taskList.Tasks[i].UpdatedAt = time.Now().UTC()
			s.storage.SaveTasks(taskList)
			return nil
		}
	}

	return ErrTaskNotFound
}

// method delete task
func (s *TaskService) DeleteTask(id int) error {
	taskList, _ := s.storage.LoadTasks()
	for i, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks = append(taskList.Tasks[i:], taskList.Tasks[i+1:]...)
			s.storage.SaveTasks(taskList)
			return nil
		}
	}

	return ErrTaskNotFound
}

func generateID(tasks []domain.Task) int {
	if len(tasks) == 0 {
		return 1
	}

	return tasks[len(tasks)-1].ID + 1
}
