package cli

import (
	"github.com/awesome-goose/goose/types"
	"github.com/awesome-goose/multi/app/shared"
)

// CliModule is the root module for the CLI platform
type CliModule struct{}

func (m *CliModule) Imports() []types.Module {
	return []types.Module{
		// Import shared modules
		&shared.SharedModule{},
		// Import CLI routes (commands)
		CLI_ROUTES,
	}
}

func (m *CliModule) Exports() []any {
	return []any{}
}

func (m *CliModule) Declarations() []any {
	return []any{
		&CliService{},
	}
}
