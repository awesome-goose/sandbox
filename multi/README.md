# Multi-Platform App Sandbox

This is an example of a multi-platform application using the goose framework.

## Overview

This app demonstrates how to run multiple platforms (API, Web, CLI) in a single application:

- **API Server**: Runs on port 8080, provides JSON API endpoints
- **Web Server**: Runs on port 3000, serves HTML pages
- **CLI**: Command-line interface for administrative tasks

## Running the Application

### Start API and Web servers concurrently

```bash
go run main.go
```

This starts:

- API server at http://localhost:8080
- Web server at http://localhost:3000

### Run CLI commands

```bash
go run main.go cli info
go run main.go cli help
go run main.go cli version
```

## API Endpoints

- `GET /api/status` - Get API status
- `GET /api/health` - Health check

## Web Routes

- `GET /` - Home page
- `GET /status` - Web status page

## Project Structure

```
multi/
├── main.go                 # Application entry point
├── go.mod                  # Go module file
├── .env                    # Environment configuration
└── app/
    ├── api/                    # API platform
    │   ├── api.module.go       # API platform module
    │   ├── api.controller.go   # API request handlers
    │   ├── api.service.go      # API business logic
    │   └── api.routes.go       # API route definitions
    ├── web/                    # Web platform
    │   ├── web.module.go       # Web platform module
    │   ├── web.controller.go   # Web request handlers
    │   ├── web.service.go      # Web business logic
    │   └── web.routes.go       # Web route definitions
    ├── cli/                    # CLI platform
    │   ├── cli.module.go       # CLI platform module
    │   ├── cli.controller.go   # CLI command handlers
    │   ├── cli.service.go      # CLI business logic
    │   └── cli.routes.go       # CLI command definitions
    └── shared/                 # Shared across all platforms
        ├── shared.module.go    # Shared module
        └── shared.service.go   # Shared business logic
```

## Key Concepts

1. **Shared Code**: The `shared/` directory contains code that can be used by all platforms.

2. **Platform-Specific Code**: Each platform has its own module, controller, service, and routes.

3. **Concurrent Execution**: API and Web platforms run in separate goroutines.

4. **CLI Invocation**: CLI is only executed when the `cli` argument is passed.

5. **Graceful Shutdown**: All servers handle SIGINT/SIGTERM for graceful shutdown.
