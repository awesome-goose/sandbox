package app

import (
	"github.com/awesome-goose/goose/types"
	"github.com/awesome-goose/spa/app/user"
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
