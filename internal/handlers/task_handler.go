package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"taskapi/internal/models"
	"taskapi/internal/storage"
	"taskapi/internals/models"
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