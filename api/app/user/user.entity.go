package user

import (
	"fmt"

	"github.com/awesome-goose/goose/modules/sql"
)

type User struct {
	sql.BaseEntity

	Name string `json:"name"`
}

type UserEntity struct {
	*sql.Entity[User]

	hooks map[string]func(*User) error
}

func (ue *UserEntity) OnRegister() {
	ue.hooks = map[string]func(*User) error{
		sql.BeforeCreate: func(user *User) error {
			// Example hook: Validate user data before creation
			if user.Name == "" {
				return fmt.Errorf("user name cannot be empty")
			}
			return nil
		},
	}
}
