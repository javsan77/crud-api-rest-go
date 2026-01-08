package storage

import (
	"sync"
	"taskapi/internal/models"
	"time"
)

// TaskStore defines the interface to storage operations
type TaskStore interface{
	Create(task *models.Task) error
	GetByID(id int) (*models.Task, error)
	GetAll() ([]*models.Task, error)
	Update(id int, updates *models.UpdateTaskRequest) (*models.Task, error)
	Delete(id int) error
}

// MemoryStore implements TaskStore using memory
type MemoryStore struct {
	mu sync.RWMutex
	tasks map[int]*models.Task
	nextID int
}

// NewMemoryStore creates a new instance of MemoryStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks: make(map[int]*models.Task),
		nextID:1,
	}
}

// Creates a new task
func (s *MemoryStore) Create(task*models.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = s.nextID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	s.tasks[task.ID] = task
	s.nextID++

	return nil
}

// GetByID get a task by ID
func (s *MemoryStore) GetByID(id int) (*models.Task, error){
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	if !exists {
		return  nil, models.ErrTaskNotFound
	}

	return task, nil

}

// GetAll get all tasks
func (s *MemoryStore) GetAll() ([]*models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks :=make[]*models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Update a pending task
func (s *MemoryStore) Update (id int, updates *models.UpdateTaskRequest) (*models.Task, error){
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exits {
		return nil, models.ErrTaskNotFound
	}

	// Update only the fields provided
	if updates.Title != nil {
		task.Title = *updates.Title
	}
	if updates.Description != nil {
		task.Description = *updates.Description
	}
	if updates.Completed != nil {
		task.Completed = *updates.Completed
	}

	task.Completed = time.Now()

	return task, nil
}

// Delete a task
func (s *MemoryStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return models.ErrTaskNotFound
	}

	delete(s.tasks, id)
	return nil
}