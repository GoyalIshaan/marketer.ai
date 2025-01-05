package routes

import (
	"github.com/gofiber/fiber/v2"
)

func MainRouter(app *fiber.App) {
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.SendString("Hello, users!")
	})

	app.Post("/login", AuthRouter)
}
