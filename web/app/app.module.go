package app

import (
	"github.com/awesome-goose/goose/modules/sql"
	"github.com/awesome-goose/goose/types"
	"github.com/awesome-goose/web/app/user"
)

type AppModule struct{}

func (m *AppModule) Imports() []types.Module {
	return []types.Module{
		sql.Root(&sql.Config{
			Dialect:  "postgres",
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "root",
			Pass:     "fullife",
			Name:     "goose_web",
			Sync:     true,
			Log:      true,
			SSLMode:  "disable",
			Schema:   "public",
			TimeZone: "UTC",
		}),

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
