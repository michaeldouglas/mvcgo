package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SecurityHeadersMiddleware() fiber.Handler {

	return func(c *fiber.Ctx) error {
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		fmt.Println("Bye 👋!")

		return c.Next()

	}

}
