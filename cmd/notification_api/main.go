package notification_api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pakut2/mandarin/pkg/notification"
)

func InitApi(server *fiber.App) {
	notificationService := notification.NewService()

	notificationEndpoint := server.Group("/notifications")
	notificationEndpoint.Post("/", CreateNotification(notificationService))
}
