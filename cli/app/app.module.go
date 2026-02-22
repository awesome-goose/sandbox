package app

import (
	"github.com/awesome-goose/cli/app/user"
	"github.com/awesome-goose/goose/types"
)

type AppModule struct{}

func (m *AppModule) Imports() []types.Module {
	return []types.Module{
		&user.UserModule{},

		ROUTES,
	}
}

func (m *AppModule) Exports() []any {
	return []any{}
}

func (m *AppModule) Declarations() []any {
	return []any{
		&AppService{},
	}
}
