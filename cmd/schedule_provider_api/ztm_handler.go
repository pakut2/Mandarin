package schedule_provider_api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pakut2/mandarin/pkg/logger"
)

// @Tags 	ztm
// @Summary	Get ZTM stop with all line numbers
// @Produce	json
// @Param 	stopId path string true "ZTM Stop ID"
// @Success	200
// @Failure	400
// @Router 	/ztm/stop/{stopId} [get]
func GetZtmStopLineNumbers(service ZtmService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		stopId := c.Params("stopId")

		if stopId == "" {
			logger.Logger.Errorf("stopId route param not provided")
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "stopId route param not provided"})
		}

		ztmStop, err := service.GetStopById(stopId)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		return c.JSON(ztmStop)
	}
}
