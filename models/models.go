package models

type Task struct {
	ID          int    `json:"id"`          // Task ID (Primary Key)
	Title       string `json:"title"`       // Task Title (Required)
	Description string `json:"description"` // Task Description (Optional)
	Completed   bool   `json:"completed"`   // Task Completion Status (Default: false)
}
