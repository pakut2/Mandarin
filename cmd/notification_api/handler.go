package notification_api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	notification_dto "github.com/pakut2/mandarin/cmd/notification_api/dto"
	"github.com/pakut2/mandarin/pkg/notification"
	"github.com/pakut2/mandarin/pkg/server"
)

// @Tags 	notifications
// @Summary	Create Notification
// @Accept 	json
// @Produce	json
// @Param 	notification body notification_dto.CreateNotificationDto true "Create Notification"
// @Success	200
// @Failure	400
// @Router 	/notifications [post]
func CreateNotification(service notification.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var createNotificationDto notification_dto.CreateNotificationDto

		if err := server.ParseBody(c, &createNotificationDto); err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}

		notification, err := service.CreateNotification(&createNotificationDto)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		return c.JSON(notification)
	}
}
