package application

import (
	"errors"
	"fmt"
	"strings"
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
func (s *TaskService) AddTask(description, dueDateStr, priority string) (domain.Task, error) {
	if description == "" {
		return domain.Task{}, ErrEmptyDescription
	}

	taskList, err := s.storage.LoadTasks()
	if err != nil {
		return domain.Task{}, err
	}

	// Validasi prioritas
	validPriorities := map[string]bool{
		"high":   true,
		"medium": true,
		"low":    true,
	}
	if priority != "" && !validPriorities[priority] {
		return domain.Task{}, fmt.Errorf("prioritas tidak valid (harus high, medium, low)")
	}

	var dueDate *time.Time
	if dueDateStr != "" {
		parseDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			return domain.Task{}, fmt.Errorf("format tanggal tidak valid (harus YYYY-MM-DD)")
		}
		dueDate = &parseDate
	}
	newTask := domain.Task{
		ID:          generateID(taskList.Tasks),
		Description: description,
		Status:      "todo",
		Priority:    priority,
		DueDate:     dueDate,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	taskList.Tasks = append(taskList.Tasks, newTask)
	if err := s.storage.SaveTasks(taskList); err != nil {
		return domain.Task{}, err
	}
	return newTask, nil
}

// method update task
func (s *TaskService) UpdateTaskWithDueDate(id int, description string, dueDateStr string) error {
	taskList, err := s.storage.LoadTasks()
	if err != nil {
		return err
	}

	var dueDate *time.Time
	if dueDateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			return fmt.Errorf("format tanggal tidak valid (harus YYYY-MM-DD)")
		}
		dueDate = &parsedDate
	}

	for i, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks[i].Description = description
			if dueDate != nil || dueDateStr == "" { // Jika dueDateStr kosong, hapus deadline
				taskList.Tasks[i].DueDate = dueDate
			}
			taskList.Tasks[i].UpdatedAt = time.Now().UTC()
			return s.storage.SaveTasks(taskList)
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
			return s.storage.SaveTasks(taskList)
		}
	}

	return ErrTaskNotFound
}

// method update status task
func (s *TaskService) UpdateTaskStatus(id int, status string) error {
	validStatus := map[string]bool{
		"todo":        true,
		"in-progress": true,
		"done":        true,
	}

	if !validStatus[status] {
		return ErrInvalidStatus
	}

	taskList, _ := s.storage.LoadTasks()
	for i, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks[i].Status = status
			taskList.Tasks[i].UpdatedAt = time.Now().UTC()
			return s.storage.SaveTasks(taskList)
		}
	}

	return ErrTaskNotFound
}

// method filter task
func (s *TaskService) GetTasks(filterStatus string) ([]domain.Task, error) {
	taskList, err := s.storage.LoadTasks()
	if err != nil {
		return nil, err
	}

	var filteredTasks []domain.Task
	for _, task := range taskList.Tasks {
		if filterStatus == "" || task.Status == filterStatus {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks, nil
}

// method pencarian task
func (s *TaskService) SearchTasks(keyword string) ([]domain.Task, error) {
	taskList, err := s.storage.LoadTasks()
	if err != nil {
		return nil, err
	}

	var results []domain.Task
	lowerKeyword := strings.ToLower(keyword)
	for _, task := range taskList.Tasks {
		if strings.Contains(strings.ToLower(task.Description), lowerKeyword) {
			results = append(results, task)
		}
	}

	return results, nil
}

// method filter by deadline
func (s *TaskService) GetTasksByDeadline(filter string) ([]domain.Task, error) {
	taskList, err := s.storage.LoadTasks()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	var results []domain.Task
	for _, task := range taskList.Tasks {
		if task.DueDate == nil {
			continue
		}
		switch filter {
		case "overdue":
			if task.DueDate.Before(now) {
				results = append(results, task)
			}
		case "upcoming":
			if task.DueDate.After(now) {
				results = append(results, task)
			}
		}
	}
	return results, nil
}

func generateID(tasks []domain.Task) int {
	if len(tasks) == 0 {
		return 1
	}

	return tasks[len(tasks)-1].ID + 1
}
