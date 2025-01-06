package main

import (
	"log"
	"marketer-ai-backend/database"
	"marketer-ai-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	database.ConnectDB()

	app := fiber.New()
	app.Use(logger.New())

	routes.MainRouter(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))

}