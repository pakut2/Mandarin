package main

import (
	"github.com/gofiber/swagger"
	"github.com/pakut2/mandarin/cmd/notification_api"
	"github.com/pakut2/mandarin/config"
	_ "github.com/pakut2/mandarin/docs"
	"github.com/pakut2/mandarin/pkg/database"
	"github.com/pakut2/mandarin/pkg/server"
)

func main() {
	if err := config.LoadEnvVariables(); err != nil {
		panic(err)
	}

	if err := database.InitConnection(); err != nil {
		panic(err)
	}
	defer database.CloseConnection()

	server := server.InitServer()

	notification_api.InitApi(server)

	server.Get("/api/*", swagger.HandlerDefault)
	server.Listen(":" + config.Env.PORT)
}
