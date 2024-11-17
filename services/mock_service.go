package services

import (
	"task-api/models"

	"github.com/stretchr/testify/mock"
)

// MockService is a mock implementation of the service layer
type MockService struct {
	mock.Mock
}

func (m *MockService) CreateTask(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockService) GetTasks() ([]models.Task, error) {
	args := m.Called()
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockService) GetTaskByID(id int) (models.Task, error) {
	args := m.Called(id)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockService) UpdateTask(id int, task *models.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockService) DeleteTask(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
