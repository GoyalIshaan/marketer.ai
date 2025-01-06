package routes

import (
	"marketer-ai-backend/database"
	"marketer-ai-backend/models"
	"marketer-ai-backend/validation"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// api/protected/user/me
func UserRouter(router fiber.Router) {

	userGroup := router.Group("/user")


	// @Description Get the current user
	// @Route api/protected/user/me
	userGroup.Get("/me", func(c *fiber.Ctx) error {
		userId := c.Locals("user_id").(uint)
		username := c.Locals("username").(string)
		email := c.Locals("email").(string)

		return c.JSON(fiber.Map{
			"user_id": userId,
			"username": username,
			"email": email,
		})
	})

	// @Description Get A User By ID
	// @Route api/protected/user/:id
	userGroup.Get("/:id", func(context *fiber.Ctx) error {
		userId, err := strconv.Atoi(context.Params("id"))
		if err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		if !validation.IsValidUserId(uint(userId)) {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		var user models.User
		result := database.DB.First(&user, uint(userId))
		if result.Error != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"user": user,
		})
	})

	// @Description Get All User Campaigns
	// @Route api/protected/user/:id/campaigns
	userGroup.Get("/:id/campaigns", func(context *fiber.Ctx) error {
		userId, err := strconv.Atoi(context.Params("id"))
		if err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		var campaigns []models.Campaign
		result := database.DB.Where("user_id = ?", uint(userId)).Find(&campaigns)
		if result.Error != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"campaigns": campaigns,
		})
	})
}