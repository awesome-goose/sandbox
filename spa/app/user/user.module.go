package user

import "github.com/awesome-goose/goose/types"

type UserModule struct{}

func (m *UserModule) Imports() []types.Module {
	return []types.Module{
		ROUTES,
	}
}

func (m *UserModule) Exports() []any {
	return []any{}
}

func (m *UserModule) Declarations() []any {
	return []any{
		&UserService{},
	}
}
