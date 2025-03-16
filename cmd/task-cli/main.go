package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andrMaulana/go-simple-task-tracker/infrastructure"
	"github.com/andrMaulana/go-simple-task-tracker/internal/application"
)

func main() {
	storage := infrastructure.NewJsonStorage()
	service := application.NewTaskService()

	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [args]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Deskripsi tugas tidak boleh kosong")
			return
		}
		description := os.Args[2]
		task, err := service.AddTask(description)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Tugas Berhasil Ditambahkan (ID: %d)\n", task.ID)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: ID dan deskripsi baru harus diisi")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		newDescription := os.Args[3]
		err := service.UpdateTask(id, newDescription)
		if err != nil {
			fmt.Println(err)
			return
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID harus diisi")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := service.DeleteTask(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Tugas #%d berhasil di hapus\n", id)

	case "mark-in-progress", "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID harus diisi")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		status := "in-progress"
		if command == "mark-done" {
			status = "done"
		}
		err := service.UpdateTaskStatus(id, status)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Status tugas #%d diubah ke '%s'\n", id, status)

	default:
		fmt.Println("Command tidak dikenali")
	}
}
