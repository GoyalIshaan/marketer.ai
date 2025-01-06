package routes

import (
	"marketer-ai-backend/database"
	"marketer-ai-backend/models"
	"marketer-ai-backend/validation"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CampaignRouter(router fiber.Router) {
	campaignGroup := router.Group("/campaign")
	singularCampaignGroup := campaignGroup.Group("/:id")

	ContentRoutes(singularCampaignGroup)

	campaignGroup.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, campaign!")
	})


	// @Description Create a new campaign
	// @Route api/protected/campaign/ 
	campaignGroup.Post("/", func(context *fiber.Ctx) error {
		newCampaignRequest := new(models.CampaignRequest)

		if err := context.BodyParser(newCampaignRequest); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}

		if !validation.IsValidCampaignRequest(*newCampaignRequest) {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid campaign request")
		}

		newCampaign := models.Campaign{
			Title: newCampaignRequest.Title,
			Description: newCampaignRequest.Description,
			Budget: newCampaignRequest.Budget,
			Platform: newCampaignRequest.Platform,
			Status: newCampaignRequest.Status,
			StartDate: newCampaignRequest.StartDate,
			EndDate: newCampaignRequest.EndDate,
			UserID: context.Locals("user_id").(uint),
		}

		result := database.DB.Create(&newCampaign)
		if result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to create campaign")
		}

		return context.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Campaign created successfully",
		})
	})

	// @Description Get a campaign
	// @Route api/protected/campaign/:id 
	campaignGroup.Get("/:id", func(context *fiber.Ctx) error {
		campaignId, err := strconv.Atoi(context.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid campaign ID")
		}

		if !validation.IsValidCampaignId(uint(campaignId)) {
			return fiber.NewError(fiber.StatusBadRequest, "Campaign not found")
		}

		var campaign models.Campaign
		result := database.DB.First(&campaign, uint(campaignId))
		if result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, "Campaign not found")
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"campaign": campaign,
		})
	})

	// @Description Update a campaign
	// @Route api/protected/campaign/:id 
	campaignGroup.Put("/:id", func(context *fiber.Ctx) error {
		campaignId, err := strconv.Atoi(context.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid campaign ID")
		}

		if !validation.IsValidCampaignId(uint(campaignId)) {
			return fiber.NewError(fiber.StatusBadRequest, "Campaign not found")
		}

		updateCampaignRequest := new(models.CampaignRequest)

		if err := context.BodyParser(updateCampaignRequest); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}

		if !validation.IsValidCampaignRequest(*updateCampaignRequest) {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid campaign update request")
		}

		updateCampaign := models.Campaign{
			Title: updateCampaignRequest.Title,
			Description: updateCampaignRequest.Description,
			Budget: updateCampaignRequest.Budget,
			Platform: updateCampaignRequest.Platform,
			Status: updateCampaignRequest.Status,
			StartDate: updateCampaignRequest.StartDate,
			EndDate: updateCampaignRequest.EndDate,
			ID: uint(campaignId),
			UserID: context.Locals("user_id").(uint),
		}

		result := database.DB.Model(&models.Campaign{}).Where("id = ?", updateCampaign.ID).Updates(&updateCampaign)
		if result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to update campaign")
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Campaign updated successfully",
		})
	})

	// @Description Delete a campaign
	// @Route api/protected/campaign/:id 
	campaignGroup.Delete("/:id", func(context *fiber.Ctx) error {
		campaignId, err := strconv.Atoi(context.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid campaign ID")
		}

		if !validation.IsValidCampaignId(uint(campaignId)) {
			return fiber.NewError(fiber.StatusBadRequest, "Campaign not found")
		}

		result := database.DB.Delete(&models.Campaign{}, uint(campaignId))
		if result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete campaign")
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Campaign deleted successfully",
		})
	})
}


