# CRUD - REST API in Go

A complete REST API for task management, built from scratch with pure Go (no external frameworks). An educational project ideal for learning backend development fundamentals with Golang.

## 🚀 Features

- ✅ Full CRUD operations (Create, Read, Update, Delete)
- ✅ RESTful API with well-defined endpoints
- ✅ Clean and modular architecture
- ✅ Thread-safe in-memory storage
- ✅ Custom middlewares (CORS, logging)
- ✅ Robust error handling
- ✅ No external dependencies (only Go stdlib)

## 📋 Prerequisites

- Go 1.21 or higher installed
- curl or Postman to test the API
- Operating System: Linux, macOS, or Windows

## 🛠️ Installation and Setup

### 1. Clone the repository

```bash
git clone https://github.com/javsan77/crud-api-rest-go
cd crud-api-rest-go
```

### 2. Verify Go installation

```bash
go version
```

### 3. Initialize Go modules

```bash
go mod tidy
```

### 4. Run the server

```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080`

You should see:
```
🚀 Server started on http://localhost:8080
Available endpoints:
  GET    /tasks       - Get all tasks
  POST   /tasks       - Create a task
  GET    /tasks/{id}  - Get a task
  PUT    /tasks/{id}  - Update a task
  DELETE /tasks/{id}  - Delete a task
```

## 📂 Project Structure

```
crud-api-rest-go/
├── cmd/
│   └── api/
│       └── main.go              # Entry point and HTTP server
├── internal/
│   ├── handlers/
│   │   └── task_handler.go      # HTTP handlers
│   ├── models/
│   │   └── task.go              # Data models and DTOs
│   └── storage/
│       └── memory_store.go      # Persistence layer
├── .gitignore
├── go.mod
├── LICENSE
└── README.md
```

### Component Description

- **cmd/api**: Server configuration and routing
- **internal/handlers**: HTTP request handling logic
- **internal/models**: Data structures and validations
- **internal/storage**: Storage implementation (currently in-memory)

## 🔌 API Endpoints

### 📋 Data Model: Task

```json
{
  "id": 1,
  "title": "Task title",
  "description": "Detailed description",
  "completed": false,
  "created_at": "2026-01-08T10:00:00Z",
  "updated_at": "2026-01-08T10:00:00Z"
}
```

### Available Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/tasks` | Create a new task |
| GET | `/tasks` | Get all tasks |
| GET | `/tasks/{id}` | Get task by ID |
| PUT | `/tasks/{id}` | Update a task |
| DELETE | `/tasks/{id}` | Delete a task |

## 📝 Usage Examples

### 1. Create a new task

**Request:**
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Complete REST API tutorial"
  }'
```

**Response:** `201 Created`
```json
{
  "id": 1,
  "title": "Learn Go",
  "description": "Complete REST API tutorial",
  "completed": false,
  "created_at": "2026-01-08T10:00:00Z",
  "updated_at": "2026-01-08T10:00:00Z"
}
```

### 2. Get all tasks

**Request:**
```bash
curl http://localhost:8080/tasks
```

**Response:** `200 OK`
```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "description": "Complete REST API tutorial",
    "completed": false,
    "created_at": "2026-01-08T10:00:00Z",
    "updated_at": "2026-01-08T10:00:00Z"
  }
]
```

### 3. Get a specific task

**Request:**
```bash
curl http://localhost:8080/tasks/1
```

**Response:** `200 OK`
```json
{
  "id": 1,
  "title": "Learn Go",
  "description": "Complete REST API tutorial",
  "completed": false,
  "created_at": "2026-01-08T10:00:00Z",
  "updated_at": "2026-01-08T10:00:00Z"
}
```

### 4. Update a task

**Request:**
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "completed": true
  }'
```

**Response:** `200 OK`
```json
{
  "id": 1,
  "title": "Learn Go",
  "description": "Complete REST API tutorial",
  "completed": true,
  "created_at": "2026-01-08T10:00:00Z",
  "updated_at": "2026-01-08T10:30:00Z"
}
```

### 5. Delete a task

**Request:**
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

**Response:** `204 No Content`

### Error Handling

**Task not found:**
```json
{
  "error": "Task not found"
}
```

**Validation failed:**
```json
{
  "error": "title is required"
}
```

## 🏗️ Go Concepts Implemented

This project is excellent for learning Go because it implements:

### 1. **Modular Architecture**
Clear separation of concerns in layers (handlers, models, storage)

### 2. **Interfaces**
```go
type TaskStore interface {
    Create(task *Task) error
    GetByID(id int) (*Task, error)
    // ...
}
```

### 3. **Pointers and Optional Values**
```go
type UpdateTaskRequest struct {
    Title *string `json:"title,omitempty"`  // Optional field
}
```

### 4. **Safe Concurrency**
```go
type MemoryStore struct {
    mu sync.RWMutex  // Protects concurrent access
    tasks map[int]*Task
}
```

### 5. **Middleware Pattern**
```go
handler := loggingMiddleware(corsMiddleware(mux))
```

### 6. **Explicit Error Handling**
No exceptions, all errors are handled explicitly

### 7. **Struct Tags**
```go
type Task struct {
    ID int `json:"id"`  // JSON serialization
}
```

## 🧪 Building the Application

### Build for your OS
```bash
go build -o task-api cmd/api/main.go
./task-api
```

### Cross-platform compilation
```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o task-api-linux cmd/api/main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o task-api.exe cmd/api/main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o task-api-macos cmd/api/main.go
```

## 🔄 Roadmap and Future Improvements

- [ ] **Persistence**: PostgreSQL/MySQL integration
- [ ] **Authentication**: User system and JWT
- [ ] **Tests**: Unit and integration tests
- [ ] **Advanced Validation**: Use libraries like `validator`
- [ ] **Pagination**: Support for large datasets
- [ ] **Filters and Search**: Advanced query parameters
- [ ] **Swagger/OpenAPI**: Automatic API documentation
- [ ] **Docker**: Application containerization
- [ ] **CI/CD**: Continuous integration pipelines
- [ ] **Structured Logging**: Implement `zap` or `logrus`
- [ ] **Configuration**: Environment variables with `viper`
- [ ] **Rate Limiting**: Request rate control
- [ ] **Caching**: Redis for performance improvement

## 📚 Learning Resources

- [Official Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://go.dev/tour/)

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Your Name**
- GitHub: [@javsan77](https://github.com/javsan77)
- LinkedIn: [my-profile](https://www.linkedin.com/in/javier-sanchez-ayte/)




