package server

import (
	"github.com/gofiber/fiber/v2"
)

func ParseBody(c *fiber.Ctx, body interface{}) error {
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	if err := Validate.Struct(body); err != nil {
		return err
	}

	return nil
}
