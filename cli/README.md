<p align="center">
  <strong>🪿 Goose Framework</strong><br>
  <sub>Modular • Scalable • Multi-Platform</sub>
</p>

---

# CLI Sandbox

A command-line interface application example built with the Goose framework.

---

## Features

- Command routing
- Structured command handlers
- User module with CLI commands
- Dependency injection
- Console output formatting

---

## Running

```bash
go mod tidy
go run main.go
```

---

## Commands

### App Commands

| Command | Description                    |
| ------- | ------------------------------ |
| `/`     | Health check / default command |

### User Commands

| Command     | Description    |
| ----------- | -------------- |
| `user`      | List users     |
| `user <id>` | Get user by ID |

---

## Project Structure

```
cli/
├── main.go              # Entry point, platform config
├── .env                 # Environment variables
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
cli.NewPlatform(
    cli.WithName("cli-example"),
)
```

### Environment Variables

Configure via `.env` file or system environment variables.

---

## Example Usage

```bash
# Run default command
go run main.go

# Run with arguments
go run main.go user
go run main.go user 1
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
