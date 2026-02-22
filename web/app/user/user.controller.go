package user

import (
	"github.com/awesome-goose/goose/output"
	"github.com/awesome-goose/goose/types"
)

type UserController struct {
	userService *UserService `inject:""`
	log         types.Log    `inject:""`
}

func (sc *UserController) Get(body *GetDto) types.Output {
	sc.log.Info("UserController: Get check called")
	user := sc.userService.Get()

	// Return JSON for API-style single resource
	return output.JSON(user)
}

func (sc *UserController) List(body *ListDto) types.Output {
	sc.log.Info("UserController: List check called")
	user := sc.userService.List()

	return output.View("pages/list.html", map[string]any{
		"Title": "Users",
		"Users": []map[string]any{
			{"ID": 1, "Name": user.Name},
		},
		"Year": 2026,
	}, output.WithBaseDir("app/user/templates"))
}
