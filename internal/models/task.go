package models

import (
	"time"

	"github.com/docker/docker/libnetwork/drivers/null"
)

// Task represents a task in our system
type Task struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateTaskRequest represents the data to create a task
type CreateTaskRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

// UpdateTaskRequest represents the data to update a task
type UpdateTaskRequest struct {
	Title *string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Completed string `json:"completed,omitempty"`
}

// Validate data of creation
func (r *CreateTaskRequest) Validate() error {
	if r.Title == ""{
		return ErrTitleRequired
	}
	return nil
}

// Custom Errors
type AppError struct {
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrTitleRequired = &AppError{"title is required"}
	ErrTaskNotFound = &AppError{"task not found"}
)