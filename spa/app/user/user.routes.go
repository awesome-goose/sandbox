package user

import "github.com/awesome-goose/goose/modules/router"

// Served under the platform's /api prefix:
// GET /api/users, GET /api/users/:id, POST /api/users
var (
	ROUTES = router.ForRoutes(
		router.Get("/users", []any{UserController{}, "List"}),
		router.Get("/users/:id", []any{UserController{}, "Get"}),
		router.Post("/users", []any{UserController{}, "Create"}),
	)
)
