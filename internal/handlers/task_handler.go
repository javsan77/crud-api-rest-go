package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"taskapi/internal/models"
	"taskapi/internal/storage"
)

// TashHandler maneja las peticiones HTTP para tasks
type TaskHandler struct {
	store storage.TaskStore
}

// NewTaskHandler creates a new Handler
func NewTaskHandler(store storage.TaskStore) *TaskHandler {
	return &TaskHandler{store: store}
}

// CreateTask manages POST /tasks
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request Body")
		return
	}

	if err := req.Validate(); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	task := &models.Task{
		Title:			req.Title,
		Description: 	req.Description,
		Completed: 		false,
	}

	if err := h.store.Create(task); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create task")
		return
	}

	respondJSON(w, http.StatusCreated, task)
}

// GetTask manages GET /tasks/{id}
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request){
	id, err := extractID(r.URL.Path)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, err := h.store.GetByID(id)
	if err != nil{
		if err == models.ErrTaskNotFound{
			respondError(w, http.StatusNotFound, "Task not found")
			return
		}
		respondError(w, http.StatusInternalServerError, "Failed to get task")
		return
	}

	respondJSON(w, http.StatusOK, task)
}

// GetAllTasks maneja GET /tasks
func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request){
	tasks, err := h.store.GetAll()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to get tasks")
		return
	}

	respondJSON(w, http.StatusOK, tasks)
}

// UpdateTask manages PUT /tasks/{id}
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request){
	id, err := extractID(r.URL.Path)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	var req models.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request Body")
		return
	}

	task, err := h.store.Update(id, &req)
	if err != nil {
		if err == models.ErrTaskNotFound {
			respondError(w, http.StatusNotFound, "Task not found")
			return
		}
		respondError(w, http.StatusInternalServerError, "Failed to update task")
		return
	}
	
	respondJSON(w, http.StatusOK, task)
}


// DeleteTask manages DELETE /tasks/{id}
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	err = h.store.Delete(id)
	if err != nil {
		if err == models.ErrTaskNotFound {
			respondError(w, http.StatusNotFound, "Task not found")
			return
		}
		respondError(w, http.StatusInternalServerError, "Failed to delete task")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Auxiliary Functions
func respondJSON(w http.ResponseWriter, status int, data interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string){
	respondJSON(w, status, map[string]string{"error": message})
}

func extractID(path string) (int, error){
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 {
		return 0, models.ErrTaskNotFound
	}
	return strconv.Atoi(parts[len(parts)-1])
}