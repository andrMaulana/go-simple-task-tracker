package application

import (
	"errors"

	"github.com/andrMaulana/go-simple-task-tracker/infrastructure"
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
