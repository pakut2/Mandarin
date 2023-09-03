package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pakut2/mandarin/config"
)

func InitServer() *fiber.App {
	server := fiber.New()

	server.Use(recover.New())
	server.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return config.Env.GO_ENV == "development"
		},
	}))
	server.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	initValidator()

	return server
}
