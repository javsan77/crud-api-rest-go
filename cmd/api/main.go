package api

import (
	"fmt"
	"log"
	"net/http"
	"taskapi/internal/handlers"
	"taskapi/internal/storage"
)

func main() {
	// Initialize storage
	store := storage.NewMemoryStore()

	// Create handler
	taskHandler := handlers.NewTaskHandler(store)
	
	// Configure router
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetAllTasks(w, r)
		case http.MethodPost:
			taskHandler.CreateTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetTask(w, r)
		case http.MethodPut:
			taskHandler.UpdateTask(w, r)
		case http.MethodDelete:
			taskHandler.DeleteTask(w, r)			
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)			
		}
	})

	// Middleware para logging and CORS
	handler := loggingMiddleware(corsMiddleware(mux))

	// Start server
	port := ":8080"
	fmt.Printf("🚀 Servidor iniciado en http://localhost%s\n", port)
	fmt.Println("Endpoints disponibles:")
	fmt.Println("  GET    /tasks       - Obtener todas las tareas")
	fmt.Println("  POST   /tasks       - Crear una tarea")
	fmt.Println("  GET    /tasks/{id}  - Obtener una tarea")
	fmt.Println("  PUT    /tasks/{id}  - Actualizar una tarea")
	fmt.Println("  DELETE /tasks/{id}  - Eliminar una tarea")

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal(err)
	}

}

// loggingMiddleware registra cada petición
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// corsMiddleware habilita CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})

}