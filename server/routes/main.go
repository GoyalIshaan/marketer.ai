package routes

import (
	"marketer-ai-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func MainRouter(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/login", LoginRouter)
	api.Post("/register", RegisterRouter)

	protectedRoutes := api.Group("/protected", middleware.Protected())

	// @Description Check if user has access
	// @Route api/protected/checkaccess
	protectedRoutes.Get("/checkaccess", func (c *fiber.Ctx) error {
		return c.SendString("access granted")
	})

	UserRouter(protectedRoutes)
	CampaignRouter(protectedRoutes)
}
