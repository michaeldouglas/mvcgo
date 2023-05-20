package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/template/html"
	"github.com/michaeldouglas/mvcgo/commands"
	"github.com/michaeldouglas/mvcgo/initializers"
	"github.com/michaeldouglas/mvcgo/middleware"
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
		// Carrega a engine
		engine := html.New("./views", ".html")

		// Configuração do Fiber
		app := fiber.New(fiber.Config{
			Views:       engine,
			ViewsLayout: "layouts/main",
		})

		// Configura os arquivos estáticos
		app.Static("/", "./public")

		// Configura o middleware
		app.Use(middleware.SecurityHeadersMiddleware())
		app.Use(compress.New())

		// Nossas Rotas
		Routes(app)

		// Inicia nossa aplicação
		app.Listen(":" + os.Getenv("WEB_PORT"))
	}
}
