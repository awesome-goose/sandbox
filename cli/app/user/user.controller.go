package user

import (
	"fmt"

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
	return output.Box("User Details", []string{
		fmt.Sprintf("Name: %s", user.Name),
	})
}

func (sc *UserController) List(body *ListDto) types.Output {
	sc.log.Info("UserController: List check called")
	user := sc.userService.List()
	return output.Table(
		[]string{"Name"},
		[][]string{{user.Name}},
	)
}
