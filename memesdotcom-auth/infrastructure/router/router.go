package router

import (
	"memesdotcom-auth/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateRestRouter(usersHandler handlers.AuthHandler) *fiber.App {
	app := fiber.New()

	users := app.Group("/auth")
	users.Use(logger.New())
	{
		users.Get("/health", func(c *fiber.Ctx) error {
			return c.SendString("ok")
		})
	}

	return app
}
