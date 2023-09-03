package schedule_provider_api

import "github.com/gofiber/fiber/v2"

func InitApi(server *fiber.App) {
	ztmStopService := NewZtmService()

	ztmEndpoint := server.Group("/ztm")
	ztmEndpoint.Get("/stop/:stopId", GetZtmStopLineNumbers(ztmStopService))
}
