package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func MainHandler(c *fiber.Ctx) error {
	header := c.Request().Header.Peek("teste")
	return c.Status(200).SendString(string(header))
}
