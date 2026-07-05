package user

import (
	"strconv"

	"github.com/awesome-goose/goose/io/output"
	"github.com/awesome-goose/goose/types"
)

type UserController struct {
	userService *UserService `inject:""`
	log         types.Log    `inject:""`
}

func (c *UserController) List(body *ListDto) types.Output {
	c.log.Info("UserController: List called")

	return output.JSON(c.userService.List())
}

func (c *UserController) Get(body *GetDto) types.Output {
	c.log.Info("UserController: Get called")

	id, err := strconv.Atoi(body.ID)
	if err != nil {
		return output.JSON(map[string]any{"error": "invalid user id"})
	}

	user, found := c.userService.Get(id)
	if !found {
		return output.JSON(map[string]any{"error": "user not found"})
	}

	return output.JSON(user)
}

func (c *UserController) Create(body *CreateDto) types.Output {
	c.log.Info("UserController: Create called")

	if body.Name == "" || body.Email == "" {
		return output.JSON(map[string]any{"error": "name and email are required"})
	}

	return output.JSON(c.userService.Create(body.Name, body.Email))
}
