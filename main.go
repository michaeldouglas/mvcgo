package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/michaeldouglas/mvcgo/commands"
	"github.com/michaeldouglas/mvcgo/controllers"
	"github.com/michaeldouglas/mvcgo/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.Connection()
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		commands.Run()
	} else {
		// Configuração do Fiber
		app := fiber.New()
		// Nossas Rotas
		app.Get("/", controllers.HelloIndex)
		app.Get("/json", controllers.HelloJson)

		// Inicia nossa aplicação
		app.Listen(":" + os.Getenv("WEB_PORT"))
	}
}
