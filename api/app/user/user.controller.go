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
	return output.JSON(sc.userService.Get())
}

func (sc *UserController) List(body *ListDto) types.Output {
	sc.log.Info("UserController: List check called")
	return output.JSON(sc.userService.List())
}
