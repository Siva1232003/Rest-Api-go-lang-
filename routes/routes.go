package routes

import (
    "net/http"

    "github.com/gorilla/mux"
    "task-api/handlers"
)

func RegisterRoutes() *mux.Router {
    // Initialize a new router
    router := mux.NewRouter()

    // Task Routes
	
    router.HandleFunc("/tasks", handlers.GetTasks).Methods(http.MethodGet)   // Fetch all tasks
    router.HandleFunc("/tasks/{id}", handlers.GetTaskByID).Methods(http.MethodGet) // Fetch a task by ID
    router.HandleFunc("/tasks", handlers.CreateTask).Methods(http.MethodPost)  // Create a new task
    router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods(http.MethodPut) // Update a task
    router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods(http.MethodDelete) // Delete a task

    return router
}
