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
	initializers.SyncDB()
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		commands.Run(args)
	} else {
		// Configuração do Fiber
		app := fiber.New()
		// Nossas Rotas
		app.Get("/", controllers.HelloIndex)
		app.Get("/json", controllers.HelloJson)
		app.Get("/teste", controllers.Teste)

		// Inicia nossa aplicação
		app.Listen(":" + os.Getenv("WEB_PORT"))
	}
}
