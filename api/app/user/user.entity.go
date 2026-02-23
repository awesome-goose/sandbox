package user

import (
	"github.com/awesome-goose/goose/modules/sql"
)

type User struct {
	sql.BaseEntity

	Name string `json:"name"`
}

type UserEntity struct {
	*sql.Entity[User] `inject:""`
}

func (ue *UserEntity) OnRegister() {
	ue.Hydrate(
		"user",
		[]string{},
		[]string{},
		nil,
		nil,
		nil,
		nil,
		"",
	)
}
