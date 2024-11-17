package services

import (
	"database/sql"
	// "net/http"
	"fmt"
	"log"
	"task-api/db"
	"task-api/models"
)

// CreateTask creates a new task in the database.
func CreateTask(task *models.Task) error {
	// Check if task with the same title already exists

	if len(task.Title) > 25 {
		return fmt.Errorf("title cannot exceed 25 characters")
	}else if task.Title == "" {
		return fmt.Errorf("title should not be empty")
        
    }

	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM tasks WHERE title=$1)", task.Title).Scan(&exists)
	if err != nil {
		log.Print("welcome", err)
		return fmt.Errorf("error checking title existence: %w", err) // Wrap error for debugging
	}
	if exists {
		return fmt.Errorf("title already exists")
	}

	// Insert the new task
	err = db.DB.QueryRow("INSERT INTO tasks (title, description, completed) VALUES ($1, $2, $3) RETURNING id", task.Title, task.Description, task.Completed).Scan(&task.ID)
	if err != nil {
		return fmt.Errorf("error inserting new task: %w", err)
	}
	return nil

}

// GetTasks retrieves all tasks from the database.
func GetTasks() ([]models.Task, error) {
	rows, err := db.DB.Query("SELECT id, title, description, completed FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks: %w", err)
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil {
			fmt.Printf("error closing rows: %v\n", cerr)
		}
	}()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			return nil, fmt.Errorf("error scanning task: %w", err)
		}
		tasks = append(tasks, task)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return tasks, nil
}

// GetTaskByID retrieves a task by ID from the database.
func GetTaskByID(id int) (*models.Task, error) {
	var task models.Task
	err := db.DB.QueryRow("SELECT id, title, description, completed FROM tasks WHERE id=$1", id).Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error retrieving task by ID: %w", err)
	}
	return &task, nil
}

// UpdateTask updates the task details in the database.
func UpdateTask(id int, task *models.Task) error {
	_, err := db.DB.Exec("UPDATE tasks SET title=$1, description=$2, completed=$3 WHERE id=$4", task.Title, task.Description, task.Completed, id)
	if err != nil {
		return fmt.Errorf("error updating task: %w", err)
	}
	return nil
}

// DeleteTask deletes a task from the database.
func DeleteTask(id int) error {
	// Execute the DELETE query
	result, err := db.DB.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}

	// Check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no task found with ID: %d", id)
	}

	// Log the successful deletion
	log.Printf("Task with ID %d was successfully deleted", id)

	return nil
}
