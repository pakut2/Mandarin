package main

import (
	"github.com/pakut2/mandarin/config"
	"github.com/pakut2/mandarin/pkg/database"
	"github.com/pakut2/mandarin/pkg/server"
)

func main() {
	err := config.LoadEnvVariables()
	if err != nil {
		panic(err)
	}

	err = database.InitConnection()
	if err != nil {
		panic(err)
	}

	defer database.CloseConnection()

	server := server.InitServer()
	if err != nil {
		panic(err)
	}

	server.Listen(":" + config.Env.PORT)
}
