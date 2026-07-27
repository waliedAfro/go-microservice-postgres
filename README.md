# Go REST API with PostgreSQL

A  REST API built with Go, PostgreSQL, and the Repository-Service-Handler architecture.

This project demonstrates how to build a clean and maintainable backend using dependency injection, layered architecture, and PostgreSQL connection pooling.

---

# Features

- RESTful API
- PostgreSQL Database
- Repository Pattern
- Service Layer
- HTTP Handlers
- Dependency Injection
- Connection Pooling (pgxpool)
- Environment Configuration (.env)
- Logging Middleware
- Graceful HTTP Server Configuration
- Pagination

---

# Project Structure

```
internal 
├── config/
│   └── config.go
├── database/
│   └── postgres.go
├── handler/
│   └── book_handler.go
├── middleware/
│   └── logging.go
├── models/
│   └── book.go
├── dto/
│      create_book_request.go
│      update_book_request.go
│      book_filter.go
│
├── validation/
│      validator.go
├── repository/
│   └── postgres_book_repository.go
    └──postgres_book_repository.go
    └──helper
        └──generate_query.go
├── service/
│   └── bool_service_imp.go
├       book_service.go
├── common
├   └──pagination.go
|   
├── main.go
├── go.mod
├── .env
└── README.md
```

---

# Architecture

```
                HTTP Request
                      │
                      ▼
                HTTP Handler
                      │
                      ▼
                 Service Layer
                      │
                      ▼
              Repository Layer
                      │
                      ▼
                 PostgreSQL
```

Each layer has a single responsibility.

| Layer | Responsibility |
|--------|---------------|
| Handler | HTTP requests and responses |
| Service | Business logic |
| Repository | Database operations |
| Database | PostgreSQL connection |

---

# Technologies

- Go 1.24+
- PostgreSQL
- pgx v5
- pgxpool
- net/http
- godotenv

---

# Installation

Clone the repository.

```bash
git clone https://github.com/waliedAfro/go-microservice-postgres.git

cd go-microservice-postgres
```

---

Install dependencies.

```bash
go mod tidy
```

---

# Database

Create a PostgreSQL database.

Example:

```sql
CREATE DATABASE books_db;
```

Create the table.

```sql
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL
);
```

---

# Environment Variables

Create a `.env` file in the project root.

```env
GO_DATABASE_URL=postgres://postgres:password@localhost:5432/books_db?sslmode=disable
```

Load the environment variables using `godotenv`.

```go
godotenv.Load()
```

---

# Running the Application

```bash
go run ./cmd/api .
```

Output:

```
Connected to PostgreSQL
Server started on :8080
```

---

# API Endpoints

## Get all books

```
GET /books
```

Response

```json
[
    {
        "id":1,
        "title":"Clean Code",
        "author":"Robert C. Martin"
    }
]
```

---

## Get book by ID

```
GET /books/{id}
```

Example

```
GET /books/1
```

---

## Create a book

```
POST /books
```

Request

```json
{
    "title":"Go in Action",
    "author":"William Kennedy"
}
```

Response

```json
{
    "id":1,
    "title":"Go in Action",
    "author":"William Kennedy"
}
```

---

## Delete a book

```
DELETE /books/{id}
```

Example

```
DELETE /books/1
```

Response

```
204 No Content
```

---

# PostgreSQL Connection Pool

The application uses `pgxpool`.

Example configuration:

```go
cfg.MaxConns = 5
cfg.MinConns = 2
cfg.MaxConnLifetime = time.Hour
cfg.MaxConnIdleTime = 30 * time.Minute
```

Benefits:

- Faster database access
- Reuses connections
- Better scalability
- Automatic connection management

---

# Dependency Injection

Dependencies are created in `main.go`.

```go
db := database.NewPostgres(cfg.DatabaseURL)

bookRepo := repository.NewPostgresBookRepository(db)

bookService := service.NewBookService(bookRepo)

bookHandler := handler.NewBookHandler(bookService)
```

This approach keeps each layer independent and testable.

---

# Error Handling

Examples:

- Invalid JSON
- Invalid Book ID
- Book Not Found
- Database Errors
- Internal Server Errors

HTTP status codes are returned appropriately.

| Code | Description |
|------|-------------|
|200|OK|
|201|Created|
|204|No Content|
|400|Bad Request|
|404|Not Found|
|500|Internal Server Error|

---

# Logging

All HTTP requests pass through a logging middleware.

Example log:

```
GET /books
POST /books
DELETE /books/1
```

---

# Best Practices Used

- Repository Pattern
- Service Layer Pattern
- Dependency Injection
- Environment Variables
- Connection Pooling
- JSON APIs
- Separation of Concerns
- Layered Architecture
- Proper HTTP Status Codes
- Clean Project Structure

---

# Future Improvements


- Filtering
- Search
- Validation
- JWT Authentication
- Swagger/OpenAPI
- Docker Support
- Unit Tests
- Integration Tests
- Database Migrations
- Structured Logging
- Configuration Management

---

# Learning Outcomes

After completing this project, you should understand how to:

- Build REST APIs using Go.
- Organize a Go project using layered architecture.
- Connect Go to PostgreSQL.
- Use `pgxpool` for efficient database connections.
- Implement the Repository pattern.
- Implement a Service layer.
- Create HTTP handlers with `net/http`.
- Configure applications using environment variables.
- Apply dependency injection.
- Return proper HTTP responses and status codes.

---

# License

This project is provided for educational purposes. Feel free to modify and extend it for your own learning or production projects.
