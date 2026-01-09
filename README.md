# Task API – Pure Go REST API

A simple and clean **RESTful API for task management**, built **from scratch using pure Go** (standard library only for HTTP).  
This project is intended as an **educational backend example**, showing how to structure a Go application without heavy frameworks.

## 🚀 Features

- ✅ Full CRUD operations (Create, Read, Update, Delete)
- ✅ RESTful API design
- ✅ Pure Go (`net/http`) – no web frameworks
- ✅ PostgreSQL integration using `pgx`
- ✅ Clean architecture (handlers, models, storage, config)
- ✅ Partial updates using `COALESCE`
- ✅ Middleware support (Logging, CORS)
- ✅ Environment-based configuration (`.env`)
- ✅ Interface-based storage (PostgreSQL or in-memory)

---

## 🧱 Tech Stack

- **Go** 1.21+
- **PostgreSQL**
- **pgx** (PostgreSQL driver)
- **godotenv** (environment variables)

---

## 📂 Project Structure

```

taskapi/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Environment configuration
│   ├── handlers/
│   │   └── task_handler.go      # HTTP handlers
│   ├── models/
│   │   └── task.go              # Domain models and DTOs
│   └── storage/
│       ├── memory_store.go      # In-memory storage (optional)
│       └── postgres_store.go    # PostgreSQL implementation
├── .env                         # Environment variables
├── go.mod
├── go.sum
├── LICENSE
└── README.md

````

---

## ⚙️ Configuration

Create a `.env` file in the project root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=taskuser
DB_PASSWORD=taskpass123
DB_NAME=taskdb
SERVER_PORT=8080
````

---

## 🗄️ Database Setup

Example table definition:

```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
```

---

## ▶️ Running the Application

### 1. Clone the repository

```bash
git clone https://github.com/javsan77/taskapi.git
cd taskapi
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the server

```bash
go run cmd/api/main.go
```

Server will start at:

```
http://localhost:8080
```

---

## 🔌 API Endpoints

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| POST   | `/tasks`      | Create a new task |
| GET    | `/tasks`      | Get all tasks     |
| GET    | `/tasks/{id}` | Get task by ID    |
| PUT    | `/tasks/{id}` | Update a task     |
| DELETE | `/tasks/{id}` | Delete a task     |

---

## 🧾 Task Model

```json
{
  "id": 1,
  "title": "Learn Go",
  "description": "Build a REST API",
  "completed": false,
  "created_at": "2026-01-09T12:00:00Z",
  "updated_at": "2026-01-09T12:00:00Z"
}
```

---

## 📌 Example Requests

### Create a Task

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Build REST API with pure Go"
  }'
```

---

### Update a Task (Partial Update)

```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "completed": true
  }'
```

---

### Delete a Task

```bash
curl -X DELETE http://localhost:8080/tasks/1
```

---

## 🧠 Design Highlights

### ✔ Interface-Based Storage

```go
type TaskStore interface {
    Create(task *Task) error
    GetByID(id int) (*Task, error)
    GetAll() ([]*Task, error)
    Update(id int, updates *UpdateTaskRequest) (*Task, error)
    Delete(id int) error
}
```

Allows easy switching between **PostgreSQL** and **in-memory** storage.

---

### ✔ Partial Updates with COALESCE

```sql
UPDATE tasks
SET title = COALESCE($1, title),
    description = COALESCE($2, description),
    completed = COALESCE($3, completed),
    updated_at = $4
WHERE id = $5;
```

Only updates fields that are provided.

---

### ✔ Middleware Pattern

```go
handler := loggingMiddleware(corsMiddleware(mux))
```

---

## 🛠 Build Binary

```bash
go build -o task-api cmd/api/main.go
./task-api
```

### Cross Compilation

```bash
GOOS=linux GOARCH=amd64 go build -o task-api-linux cmd/api/main.go
GOOS=windows GOARCH=amd64 go build -o task-api.exe cmd/api/main.go
GOOS=darwin GOARCH=amd64 go build -o task-api-macos cmd/api/main.go
```

---

## 🗺️ Roadmap

* [ ] Unit and integration tests
* [ ] Pagination and filtering
* [ ] Authentication (JWT)
* [ ] Swagger / OpenAPI
* [ ] Docker support
* [ ] Structured logging
* [ ] CI/CD pipeline

---

## 📄 License

MIT License
See the [LICENSE](LICENSE) file for details.

---

## 👤 Author

**Javier Sanchez Ayte**

* GitHub: [https://github.com/javsan77](https://github.com/javsan77)
* LinkedIn: [https://www.linkedin.com/in/javier-sanchez-ayte/](https://www.linkedin.com/in/javier-sanchez-ayte/)
