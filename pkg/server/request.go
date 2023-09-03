package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pakut2/mandarin/pkg/logger"
)

func ParseBody(c *fiber.Ctx, body interface{}) error {
	if err := c.BodyParser(&body); err != nil {
		logger.Logger.Errorf("error parsing request body, err: %v", err)
		return err
	}

	if err := Validate.Struct(body); err != nil {
		logger.Logger.Errorf("error validating request body, err: %v", err)
		return err
	}

	return nil
}
