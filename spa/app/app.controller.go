package app

import (
	"github.com/awesome-goose/goose/io/output"
	"github.com/awesome-goose/goose/types"
)

type AppController struct {
	appService *AppService `inject:""`
	log        types.Log   `inject:""`
}

// Health is declared at GET / and served at GET /api
func (c *AppController) Health(body *HealthDto) types.Output {
	c.log.Info("AppController: Health check called")

	return output.JSON(map[string]any{
		"status": c.appService.Health(),
		"query":  body.Type,
	})
}

// Version is declared at GET /version and served at GET /api/version
func (c *AppController) Version(body *VersionDto) types.Output {
	return output.JSON(map[string]any{
		"name":    "spa-example",
		"version": "0.0.1",
	})
}
