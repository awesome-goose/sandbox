package main

import (
	"os"

	"github.com/awesome-goose/goose"
	"github.com/awesome-goose/goose/log"
	"github.com/awesome-goose/goose/log/formatters"
	"github.com/awesome-goose/goose/log/modifiers"
	"github.com/awesome-goose/goose/log/processors"
	"github.com/awesome-goose/goose/platforms/web"
	"github.com/awesome-goose/goose/types"
	"github.com/awesome-goose/web/app"
)

func main() {
	stop, err := goose.Start(
		goose.Web(webPlatform, rootModule, initializers),
	)
	if err != nil {
		panic(err)
	}

	defer stop()
}

var (
	webPlatform = web.NewPlatform(
		web.WithName("web-example"),
		web.WithPort(8080),
	)
	rootModule   = &app.AppModule{}
	initializers = []func(container types.Container) error{
		func(container types.Container) error {
			return container.Register(
				func() types.Log {
					return log.NewLog(
						log.AppLogChannel("std"),
						log.NewLogger(
							[]types.Modifier{
								modifiers.NewUUID(),
								modifiers.NewColorTagsModifier(),
								modifiers.NewSystemInfo(),
								modifiers.NewStackTrace(),
							},
							formatters.NewSyslog("app", os.Getpid()),
							processors.NewConsole(),
						),
					)
				},
				"",
				true,
			)
		},
	}
)
