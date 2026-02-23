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
	sc.log.Info("AppController: Health check called", body)
	status := sc.appService.Health()
	return output.JSON(map[string]any{
		"status": status,
		"query":  body.Type,
	})
}
