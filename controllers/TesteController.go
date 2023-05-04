package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Teste(c *fiber.Ctx) error {
	return c.SendString("Hello, World! FuncName: Teste")
}
