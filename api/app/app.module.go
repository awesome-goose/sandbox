package app

import (
	"github.com/awesome-goose/api/app/user"
	"github.com/awesome-goose/goose/modules/sql"
	"github.com/awesome-goose/goose/types"
)

type AppModule struct {
	env types.Env `inject:""`
}

func (m *AppModule) Imports() []types.Module {
	return []types.Module{
		sql.Root(&sql.Config{
			Dialect:  m.env.Get("DB_DIALECT"),
			Host:     m.env.Get("DB_HOST"),
			Port:     m.env.GetInt("DB_PORT"),
			User:     m.env.Get("DB_USER"),
			Pass:     m.env.Get("DB_PASSWORD"),
			Name:     m.env.Get("DB_NAME"),
			Sync:     m.env.GetBool("DB_SYNC"),
			Log:      m.env.GetBool("DB_LOG"),
			SSLMode:  m.env.Get("DB_SSLMODE"),
			Schema:   m.env.Get("DB_SCHEMA"),
			TimeZone: m.env.Get("DB_TIMEZONE"),
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
