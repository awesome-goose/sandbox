<p align="center">
  <strong>🪿 Goose Framework</strong><br>
  <sub>Modular • Scalable • Multi-Platform</sub>
</p>

---

# Multi-Platform Sandbox

A multi-platform application example demonstrating how to run API, Web, and CLI from a single codebase using the Goose framework.

---

## Features

- **API Server** - REST API on port 8080
- **Web Server** - HTML pages on port 3000
- **CLI** - Command-line interface
- **Shared Module** - Common services across platforms
- **Concurrent Execution** - API and Web run simultaneously
- **Graceful Shutdown** - Handles SIGINT/SIGTERM

---

## Running

### Start API and Web Servers

```bash
go mod tidy
go run main.go
```

Starts:

- API server at http://localhost:8080
- Web server at http://localhost:3000

### Run CLI Commands

```bash
go run main.go cli info
go run main.go cli help
go run main.go cli version
```

---

## Endpoints

### API Routes (port 8080)

| Method | Path          | Description  |
| ------ | ------------- | ------------ |
| GET    | `/api/status` | API status   |
| GET    | `/api/health` | Health check |

### Web Routes (port 3000)

| Method | Path      | Description |
| ------ | --------- | ----------- |
| GET    | `/`       | Home page   |
| GET    | `/status` | Status page |

### CLI Commands

| Command       | Description   |
| ------------- | ------------- |
| `cli info`    | Show app info |
| `cli help`    | Show help     |
| `cli version` | Show version  |

---

## Project Structure

```
multi/
├── main.go              # Entry point, multi-platform setup
├── .env                 # Environment variables
├── go.mod               # Go module
└── app/
    ├── api/             # API platform module
    │   ├── api.module.go
    │   ├── api.controller.go
    │   ├── api.routes.go
    │   └── api.service.go
    ├── web/             # Web platform module
    │   ├── web.module.go
    │   ├── web.controller.go
    │   ├── web.routes.go
    │   └── web.service.go
    ├── cli/             # CLI platform module
    │   ├── cli.module.go
    │   ├── cli.controller.go
    │   ├── cli.routes.go
    │   └── cli.service.go
    └── shared/          # Shared across platforms
        ├── shared.module.go
        └── shared.service.go
```

---

## Configuration

### Platform Setup

```go
// API Platform
apiPlatform := api.NewPlatform(
    api.WithName("multi-api"),
    api.WithPort(8080),
)

// Web Platform
webPlatform := web.NewPlatform(
    web.WithName("multi-web"),
    web.WithPort(3000),
)

// CLI Platform
cliPlatform := cli.NewPlatform(
    cli.WithName("multi-cli"),
)

// Start all platforms
goose.Start(
    goose.API(apiPlatform, apiModule, initializers),
    goose.Web(webPlatform, webModule, initializers),
    goose.CLI(cliPlatform, cliModule, initializers),
)
```

---

## Key Concepts

1. **Shared Code** - The `shared/` directory contains services used by all platforms
2. **Platform-Specific Code** - Each platform has its own module, controller, service, and routes
3. **Concurrent Execution** - API and Web platforms run in separate goroutines
4. **CLI Invocation** - CLI only executes when `cli` argument is passed
5. **Graceful Shutdown** - All servers handle SIGINT/SIGTERM

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
