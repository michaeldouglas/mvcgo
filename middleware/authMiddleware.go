package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func AuthMiddleware() fiber.Handler {
	config := basicauth.Config{
		Users: map[string]string{
			"admin": "123456",
		},
	}

	return basicauth.New(config)

}
