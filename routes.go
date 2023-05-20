package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/michaeldouglas/mvcgo/controllers"
	"github.com/michaeldouglas/mvcgo/middleware"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.HelloIndex)

	// Routes API
	handler := func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	}
	api := app.Group("/api", middleware.AuthMiddleware())
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	v1.Get("/json", controllers.HelloJson)
	v1.Get("/teste", handler)
}
