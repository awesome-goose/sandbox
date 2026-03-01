package app

import (
	"github.com/awesome-goose/api/app/user"
	"github.com/awesome-goose/goose/modules/sql"
	"github.com/awesome-goose/goose/types"
)

type AppModule struct {
	config types.Config `inject:""`
	env    types.Env    `inject:""`
}

func (m *AppModule) Imports() []types.Module {
	return []types.Module{
		sql.Root(&sql.Config{
			Dialect:  m.config.Get("config.sql.dialect"),
			Host:     m.env.Get("DB_HOST"),
			Port:     m.env.GetInt("DB_PORT"),
			User:     m.env.Get("DB_USER"),
			Pass:     m.env.Get("DB_PASSWORD"),
			Name:     m.env.Get("DB_NAME"),
			Sync:     m.config.GetBool("config.sql.sync"),
			Log:      m.config.GetBool("config.sql.log"),
			SSLMode:  m.config.Get("config.sql.sslmode"),
			Schema:   m.config.Get("config.sql.schema"),
			TimeZone: m.config.Get("config.sql.timezone"),
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
