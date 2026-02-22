package app

import (
	"fmt"

	"github.com/awesome-goose/goose/output"
	"github.com/awesome-goose/goose/types"
)

type AppController struct {
	appService *AppService `inject:""`
	log        types.Log   `inject:""`
}

func (sc *AppController) Health(body *HealthDto) types.Output {
	sc.log.Info("AppController: Health check called", body)
	status := sc.appService.Health()
	if body.Type != "" {
		return output.ConsoleSuccess(fmt.Sprintf("%s (query: %s)", status, body.Type))
	}
	return output.ConsoleSuccess(status)
}
