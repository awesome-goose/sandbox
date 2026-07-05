<p align="center">
  <strong>🪿 Goose Framework</strong><br>
  <sub>Modular • Scalable • Multi-Platform</sub>
</p>

---

# SPA Example (Angular)

A Goose single-page application: one Go service on port 8080 that serves
JSON API routes under `/api` and the built Angular frontend for every other
path, with `index.html` fallback for client-side routing.

---

## Structure

```
spa/
├── main.go              # Bootstraps the goose SPA platform (:8080)
├── .env                 # Environment config
├── Makefile             # install / dev / build / dist workflow
├── app/
│   ├── app.module.go    # Root module
│   ├── app.controller.go # Health + Version JSON endpoints
│   ├── app.routes.go    # Routes (served under /api)
│   └── user/            # In-memory user resource (no database needed)
├── frontend/            # Angular source (builds into public/)
└── tests/               # Backend tests
```

---

## Quick Start

```bash
make install     # go mod tidy + npm install
make dev         # Go API on :8080 + Angular dev server together
```

During development the Angular dev server (http://localhost:4200) proxies
`/api` requests to the Go backend, so the frontend gets hot reload while
hitting real endpoints.

---

## API Routes

Backend routes are declared **without** the `/api` prefix — the SPA
platform serves them under it:

| Declared            | Served at              | Description             |
| ------------------- | ---------------------- | ----------------------- |
| `GET /`             | `GET /api`             | Health check            |
| `GET /version`      | `GET /api/version`     | App name + version JSON |
| `GET /users`        | `GET /api/users`       | List users (in-memory)  |
| `GET /users/:id`    | `GET /api/users/:id`   | Get one user            |
| `POST /users`       | `POST /api/users`      | Create a user           |

Any non-`/api` path serves a file from `public/` if it exists, and falls
back to `public/index.html` for client-side routes.

---

## Production

```bash
make dist                # ng build -> public/, Go binary + assets -> dist/
cd dist && ./spa-example
```

---

## Tests

```bash
go test ./tests/...
```

---

## Note on the local `replace` directive

`go.mod` contains `replace github.com/awesome-goose/goose => ../../goose`
because the `platforms/spa` package is newer than the latest published
goose release. Remove the directive once a goose version containing
`platforms/spa` is tagged.
