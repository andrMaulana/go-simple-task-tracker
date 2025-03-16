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
	service := application.NewTaskService(storage)

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
		fmt.Printf("Tugas #%d berhasil di hapus\n", id)

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

	case "list":
		var filterStatus string
		if len(os.Args) > 2 {
			filterStatus = os.Args[2]
			validStatus := map[string]bool{
				"todo":        true,
				"in-progress": true,
				"done":        true,
				"":            true, // Untuk menampilkan semua tugas
			}
			if !validStatus[filterStatus] {
				fmt.Println("Error: Status filter tidak valid")
				return
			}
		}
		tasks, err := service.GetTasks(filterStatus)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("Tidak ada tugas yang tersedia")
			return
		}

		// Format output sebagai tabel
		fmt.Printf("%-5s %-30s %-15s %-20s %-20s\n", "ID", "Deskripsi", "Status", "Dibuat", "Diperbarui")
		fmt.Println("-------------------------------------------------------------------------------------")
		for _, task := range tasks {
			createdAt := task.CreatedAt.Local().Format("2006-01-02 15:04:05")
			updatedAt := task.UpdatedAt.Local().Format("2006-01-02 15:04:05")
			fmt.Printf("%-5d %-30s %-15s %-20s %-20s\n",
				task.ID,
				truncateString(task.Description, 30),
				task.Status,
				createdAt,
				updatedAt,
			)
		}

	default:
		fmt.Println("Command tidak dikenali")
	}
}

func truncateString(s string, max int) string {
	if len(s) > max {
		return s[:max-3] + "..."
	}
	return s
}
