package app

import (
	"github.com/awesome-goose/goose/io/output"
	"github.com/awesome-goose/goose/types"
)

type AppController struct {
	appService *AppService `inject:""`
	log        types.Log   `inject:""`
}

func (sc *AppController) Health(body *HealthDto) types.Output {
	status := sc.appService.Health()
	return output.JSON(map[string]any{
		"status": status,
		"query":  body.Type,
	})
}
