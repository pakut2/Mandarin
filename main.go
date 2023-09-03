package main

import (
	"github.com/gofiber/swagger"
	"github.com/pakut2/mandarin/cmd/notification_api"
	"github.com/pakut2/mandarin/cmd/notification_scanner"
	"github.com/pakut2/mandarin/cmd/schedule_provider_api"
	"github.com/pakut2/mandarin/config"
	_ "github.com/pakut2/mandarin/docs"
	"github.com/pakut2/mandarin/pkg/database"
	firebase_admin "github.com/pakut2/mandarin/pkg/firebase"
	"github.com/pakut2/mandarin/pkg/logger"
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

	if err := logger.InitLogger(); err != nil {
		panic(err)
	}
	defer logger.Logger.Sync()

	server := server.InitServer()
	notification_api.InitApi(server)
	schedule_provider_api.InitApi(server)

	firebaseAdmin, err := firebase_admin.InitFirebaseAdmin()
	if err != nil {
		panic(err)
	}

	notification_scanner.InitScanner(firebaseAdmin)

	server.Get("/api/*", swagger.HandlerDefault)
	server.Listen(":" + config.Env.PORT)
}
