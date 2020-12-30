package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"memesdotcom-users/handlers"
)

func CreateRestRouter(usersHandler handlers.UsersHandler) *fiber.App {
	app := fiber.New()

	users := app.Group("/users")
	users.Use(logger.New())
	{
		users.Get("/health", func(c *fiber.Ctx) error {
			return c.SendString("ok")
		})
		users.Post("/register", usersHandler.CreateUser)
	}

	return app
}
