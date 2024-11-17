package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"task-api/models"
	"task-api/services"
	"testing"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// Mock services
var mockServices = services.MockService{}

func TestCreateTask(t *testing.T) {
	task := models.Task{Title: "Test Task", Description: "A sample task"}
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Call handler
	CreateTask(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var createdTask models.Task
	err := json.NewDecoder(rr.Body).Decode(&createdTask)
	assert.NoError(t, err)
	assert.Equal(t, task.Title, createdTask.Title)
}

func TestGetTasks(t *testing.T) {
	req, _ := http.NewRequest("GET", "/tasks", nil)
	rr := httptest.NewRecorder()

	// Call handler
	GetTasks(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var tasks []models.Task
	err := json.NewDecoder(rr.Body).Decode(&tasks)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(tasks), 0)
}

func TestGetTaskByID(t *testing.T) {
	taskID := 1
	req, _ := http.NewRequest("GET", "/tasks/"+strconv.Itoa(taskID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(taskID)})
	rr := httptest.NewRecorder()

	// Call handler
	GetTaskByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var task models.Task
	err := json.NewDecoder(rr.Body).Decode(&task)
	assert.NoError(t, err)
	assert.Equal(t, taskID, task.ID)
}

func TestUpdateTask(t *testing.T) {
	taskID := 1
	updatedTask := models.Task{Title: "Updated Task", Description: "Updated Description"}
	body, _ := json.Marshal(updatedTask)
	req, _ := http.NewRequest("PUT", "/tasks/"+strconv.Itoa(taskID), bytes.NewBuffer(body))
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(taskID)})
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Call handler
	UpdateTask(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteTask(t *testing.T) {
	taskID := 1
	req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(taskID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(taskID)})
	rr := httptest.NewRecorder()

	// Call handler
	DeleteTask(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
