package domain

import "time"

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`          // Akan diimplementasi di fase 3
	DueDate     *time.Time `json:"dueDate,omitempty"` // Nullable
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}
