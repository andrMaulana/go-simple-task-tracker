package application

import "errors"

var (
	ErrTaskNotFound     = errors.New("tugas tidak ditemukan")
	ErrInvalidStatus    = errors.New("status tidak valid")
	ErrEmptyDescription = errors.New("deksripksi tidak boleh kosong")
)

type TaskService struct {
	storage Storage
}

func NewTaskService(storage Storage) *TaskService {
	return &TaskService{storage: storage}
}
