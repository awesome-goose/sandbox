package app

import (
	"github.com/awesome-goose/goose/modules/router"
)

// Routes are declared without the /api prefix — the spa platform serves
// them under it (GET / -> GET /api, GET /version -> GET /api/version).
var (
	ROUTES = router.ForRoutes(
		router.Get("/", []any{AppController{}, "Health"}),
		router.Get("/version", []any{AppController{}, "Version"}),
	)
)
