package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/andrMaulana/go-simple-task-tracker/infrastructure"
	"github.com/andrMaulana/go-simple-task-tracker/internal/application"
)

const version = "1.0.0" // Versi awal

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Printf("Task Tracker CLI v%s\n", version)
		return
	}

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
		dueDate := ""
		if len(os.Args) > 3 && strings.HasPrefix(os.Args[3], "--due=") {
			dueDate = strings.TrimPrefix(os.Args[3], "--due=")
		}
		task, err := service.AddTask(description, dueDate)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Tugas Berhasil Ditambahkan (ID: %d)\n", task.ID)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: ID dan deskripsi harus diisi")
			return
		}

		// Parsing argumen
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: ID harus berupa angka")
			return
		}

		// Cari flag --due
		newDescription := ""
		dueDate := ""
		for i := 3; i < len(os.Args); i++ {
			arg := os.Args[i]
			if strings.HasPrefix(arg, "--due=") {
				dueDate = strings.TrimPrefix(arg, "--due=")
			} else {
				newDescription = arg // Ambil deskripsi baru
			}
		}

		// Validasi deskripsi tidak boleh kosong
		if newDescription == "" {
			fmt.Println("Error: Deskripsi tidak boleh kosong")
			return
		}

		// Panggil service
		err = service.UpdateTaskWithDueDate(id, newDescription, dueDate)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Tugas #%d berhasil diperbarui\n", id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID harus diisi")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := service.DeleteTask(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Tugas #%d berhasil di hapus\n", id)

	case "mark-in-progress", "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID harus diisi")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: ID harus berupa angka")
		}
		status := "in-progress"
		if command == "mark-done" {
			status = "done"
		}
		err = service.UpdateTaskStatus(id, status)
		if err != nil {
			fmt.Println("Error:", err)
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
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Error: Keyword harus diisi")
			return
		}
		keyword := os.Args[2]
		tasks, err := service.SearchTasks(keyword)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("Tidak ada tugas yang cocok dengan kata kunci")
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
