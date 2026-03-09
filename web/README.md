<p align="center">
  <strong>🪿 Goose Framework</strong><br>
  <sub>Modular • Scalable • Multi-Platform</sub>
</p>

---

# Web Sandbox

A web application example with HTML templates built with the Goose framework.

---

## Features

- HTML template rendering
- Layout-based templates (base, pages, partials)
- User module with list/show views
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

## Routes

### App Routes

| Method | Path | Description |
| ------ | ---- | ----------- |
| GET    | `/`  | Home page   |

### User Routes

| Method | Path        | Description      |
| ------ | ----------- | ---------------- |
| GET    | `/user`     | User list page   |
| GET    | `/user/:id` | User detail page |

---

## Project Structure

```
web/
├── main.go              # Entry point, platform config
├── .env                 # Environment variables
├── go.mod               # Go module
└── app/
    ├── app.module.go    # Root module
    ├── app.controller.go # App controller
    ├── app.routes.go    # Route definitions
    ├── app.service.go   # App service
    ├── app.dtos.go      # DTOs
    ├── templates/       # HTML templates
    │   ├── base/
    │   │   └── layout.html
    │   ├── pages/
    │   │   └── home.html
    │   └── partials/
    │       ├── header.html
    │       └── footer.html
    └── user/            # User module
        ├── user.module.go
        ├── user.controller.go
        ├── user.routes.go
        ├── user.service.go
        ├── user.entity.go
        ├── user.dtos.go
        └── templates/
            └── pages/
                ├── list.html
                └── show.html
```

---

## Template Structure

### Base Layout

`templates/base/layout.html` - Main layout wrapper

### Pages

`templates/pages/*.html` - Full page templates

### Partials

`templates/partials/*.html` - Reusable template fragments (header, footer)

---

## Configuration

### Platform Options

```go
web.NewPlatform(
    web.WithName("web-example"),
    web.WithPort(8080),
)
```

### Environment Variables

Configure via `.env` file or system environment variables.

---

## Example URLs

```
http://localhost:8080/          # Home page
http://localhost:8080/user      # User list
http://localhost:8080/user/1    # User detail
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
