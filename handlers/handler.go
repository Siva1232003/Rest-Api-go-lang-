package handlers

import (
    "encoding/json"
    "net/http"
	"fmt"
    "strconv"
    "task-api/models"
    "task-api/services"
    "github.com/gorilla/mux"
)

// CreateTask handles POST /tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    err := services.CreateTask(&task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Send created task with status 201
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

// GetTasks handles GET /tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := services.GetTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID handles GET /tasks/{id}
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    task, err := services.GetTaskByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
	if task.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	

    json.NewEncoder(w).Encode(task)
}

// UpdateTask handles PUT /tasks/{id}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    err = services.UpdateTask(id, &task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// DeleteTask handles DELETE /tasks/{id}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    // Extract the task ID from the URL
    idStr := mux.Vars(r)["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid task ID, must be a number", http.StatusBadRequest)
        return
    }

    // Attempt to delete the task by ID
    err = services.DeleteTask(id)
    if err != nil {
        if err.Error() == fmt.Sprintf("no task found with ID: %d", id) {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        // Internal server error for unexpected issues
        http.Error(w, "Failed to delete task", http.StatusInternalServerError)
        return
    }

    // Respond with No Content status
    // w.WriteHeader(http.StatusNoContent)
	fmt.Printf("deleted")
}
