package user

import "github.com/awesome-goose/goose/modules/router"

var (
	ROUTES = router.ForRoutes(
		router.Resource("user", UserController{}).Only("Get", "List"),
	)
)
