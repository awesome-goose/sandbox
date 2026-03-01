<p align="center">
  <strong>🪿 Goose Framework</strong><br>
  <sub>Modular • Scalable • Multi-Platform</sub>
</p>

---

# Goose Sandbox

Example applications demonstrating the Goose framework's capabilities across different platforms.

---

## Applications

| App              | Description                              | Port       |
| ---------------- | ---------------------------------------- | ---------- |
| [api/](api/)     | REST API application with JSON responses | 8080       |
| [cli/](cli/)     | Command-line interface application       | -          |
| [web/](web/)     | Web application with HTML templates      | 8080       |
| [multi/](multi/) | Multi-platform app (API + Web + CLI)     | 8080, 3000 |

---

## Quick Start

```bash
# Run the API example
cd api && go run main.go

# Run the CLI example
cd cli && go run main.go

# Run the Web example
cd web && go run main.go

# Run the Multi-platform example
cd multi && go run main.go
```

---

## Common Structure

Each sandbox app follows the standard Goose module structure:

```
app/
├── main.go              # Entry point
├── .env                 # Environment config
├── go.mod               # Go module
└── app/
    ├── app.module.go    # Root module
    ├── app.controller.go # Request handlers
    ├── app.routes.go    # Route definitions
    ├── app.service.go   # Business logic
    ├── app.dtos.go      # Data transfer objects
    └── <module>/        # Sub-modules (e.g., user/)
```

---

## Features Demonstrated

- **Module System** - Composable, importable modules
- **Dependency Injection** - Automatic service injection via `inject:""` tags
- **Routing** - Declarative route definitions
- **Resource Controllers** - CRUD operations with `router.Resource()`
- **Multiple Platforms** - API, Web, CLI from single codebase
- **Logging** - Configurable logging with formatters and processors

---

## Requirements

- Go 1.22+
- Goose framework (`github.com/awesome-goose/goose`)
