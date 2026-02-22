package user

import "github.com/awesome-goose/goose/modules/router"

var (
	ROUTES = router.ForRoutes(
		router.Cli("user", []any{UserController{}, "List"}),
		router.Cli("user/:id", []any{UserController{}, "Get"}),
	)
)
