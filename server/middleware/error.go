package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Handler {
	return func(context *fiber.Ctx) error {
		err := context.Next()
		if err != nil {
			log.Printf("Error: %v, Path: %s, Method: %s, Status Code: %d", err, context.Path(), context.Method(), context.Response().StatusCode())

			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   err.Error(),
				"status code": context.Response().StatusCode(),
				"message": "Internal Server Error",
				"code":    "INTERNAL_ERROR",
			})
		}
		return nil
	}
}
