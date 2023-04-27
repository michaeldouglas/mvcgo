package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HelloIndex(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func HelloJson(c *fiber.Ctx) error {
	data := map[string]string{
		"message": "Hello, world!",
	}

	return c.JSON(data)
}
