<p align="center">
  <strong>🪿 Goose Framework</strong><br>
  <sub>Modular • Scalable • Multi-Platform</sub>
</p>

---

# API Sandbox

A REST API application example built with the Goose framework.

---

## Features

- JSON responses
- RESTful routing
- User resource module with CRUD operations
- Dependency injection
- Structured logging

---

## Running

```bash
go mod tidy
go run main.go
```

Server starts at http://localhost:8080

---

## Endpoints

### App Routes

| Method | Path | Description  |
| ------ | ---- | ------------ |
| GET    | `/`  | Health check |

### User Routes

| Method | Path        | Description    |
| ------ | ----------- | -------------- |
| GET    | `/user`     | List all users |
| GET    | `/user/:id` | Get user by ID |

---

## Project Structure

```
api/
├── main.go              # Entry point, platform config
├── .env                 # Environment variables
├── config.yaml          # Application config
├── go.mod               # Go module
└── app/
    ├── app.module.go    # Root module
    ├── app.controller.go # App controller
    ├── app.routes.go    # Route definitions
    ├── app.service.go   # App service
    ├── app.dtos.go      # DTOs
    └── user/            # User module
        ├── user.module.go
        ├── user.controller.go
        ├── user.routes.go
        ├── user.service.go
        ├── user.entity.go
        └── user.dtos.go
```

---

## Configuration

### Platform Options

```go
api.NewPlatform(
    api.WithName("api-example"),
    api.WithPort(8080),
)
```

### Environment Variables

Configure via `.env` file or system environment variables.

---

## Example Request

```bash
# Health check
curl http://localhost:8080/

# List users
curl http://localhost:8080/user

# Get user by ID
curl http://localhost:8080/user/1
```

---

## Running Tests

```bash
# Run all tests
go test ./tests/...

# Run with verbose output
go test ./tests/... -v
```

---

## Code Coverage

```bash
# Coverage for all goose packages
go test ./tests/... -coverpkg=./...

```
