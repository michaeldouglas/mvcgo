package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HelloIndex(c *fiber.Ctx) error {
	return c.Render("hello/index", fiber.Map{
		"message": "Hello, World!",
	})
}

func HelloJson(c *fiber.Ctx) error {
	data := map[string]string{
		"message": "Hello, world!",
	}

	return c.JSON(data)
}
