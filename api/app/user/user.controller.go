package user

import (
	"github.com/awesome-goose/goose/io/resources/api"
)

type UserController struct {
	*api.Resource[User] `inject:""`

	entity *UserEntity `inject:""`
}

func (uc *UserController) OnRegister() {
	uc.Hydrate(
		uc.entity,
	)
}
