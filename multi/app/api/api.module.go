package api

import (
	"github.com/awesome-goose/goose/types"
	"github.com/awesome-goose/multi/app/shared"
)

// ApiModule is the root module for the API platform
// It can import shared modules and API-specific modules
type ApiModule struct{}

func (m *ApiModule) Imports() []types.Module {
	return []types.Module{
		// Import shared modules
		&shared.SharedModule{},
		// Import API routes
		API_ROUTES,
	}
}

func (m *ApiModule) Exports() []any {
	return []any{}
}

func (m *ApiModule) Declarations() []any {
	return []any{
		&ApiService{},
	}
}
