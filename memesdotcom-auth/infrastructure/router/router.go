package router

import (
	"memesdotcom-auth/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateRestRouter(authHandler handlers.AuthHandler) *fiber.App {
	app := fiber.New()

	auth := app.Group("/auth")
	auth.Use(logger.New())
	{
		auth.Get("/health", func(c *fiber.Ctx) error {
			return c.SendString("ok")
		})
		auth.Post("/login", authHandler.CreateAccessToken)
	}

	return app
}
