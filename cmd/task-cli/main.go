package main

import (
	"github.com/andrMaulana/go-simple-task-tracker/infrastructure"
	"github.com/andrMaulana/go-simple-task-tracker/internal/application"
)

func main() {
	storage := infrastructure.NewJsonStorage()
	service := application.NewTaskService()
}
