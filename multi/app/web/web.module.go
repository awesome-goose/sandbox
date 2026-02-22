package web

import (
	"github.com/awesome-goose/goose/types"
	"github.com/awesome-goose/multi/app/shared"
)

// WebModule is the root module for the Web platform
// It can import shared modules and Web-specific modules
type WebModule struct{}

func (m *WebModule) Imports() []types.Module {
	return []types.Module{
		// Import shared modules
		&shared.SharedModule{},
		// Import Web routes
		WEB_ROUTES,
	}
}

func (m *WebModule) Exports() []any {
	return []any{}
}

func (m *WebModule) Declarations() []any {
	return []any{
		&WebService{},
	}
}
