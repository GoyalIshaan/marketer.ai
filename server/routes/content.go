package routes

import (
	"log"
	"marketer-ai-backend/ai"
	aihandlers "marketer-ai-backend/ai/handlers"
	"marketer-ai-backend/database"
	"marketer-ai-backend/models"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ContentRoutes(router fiber.Router) {
	contentGroup := router.Group("/content")

	contentGroup.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, content!")
	})

	// @Description Generate content
	// @Route api/protected/campaign/:id/content/
	contentGroup.Post("/", func(context *fiber.Ctx) error {
		campaignId, err := strconv.Atoi(context.Params("id"))
		log.Println(context.Params("id"))
		if err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		generateContentRequest := models.GenerateContentRequest{}
		if err := context.BodyParser(&generateContentRequest); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		campaign := models.Campaign{}
		campaignFindResult := database.DB.First(&campaign, campaignId)
		if campaignFindResult.Error != nil {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": campaignFindResult.Error.Error(),
			})
		}

		promptContentResponse := models.PromptContentResponse{
			ContentType: generateContentRequest.ContentType,
			Campaign: campaign,
			Hashtags: generateContentRequest.Hashtags,
		}

		aiContent, err := ai.GenerateContent(aihandlers.PromptGenerator(promptContentResponse))
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		log.Println(aiContent)

		title , content := aihandlers.ParseContent(aiContent)
		newCampaignContent := models.Content{
			Title: title,
			Content: content,
			CampaignID: uint(campaignId),
			ContentType: generateContentRequest.ContentType,
			Hashtags: generateContentRequest.Hashtags,
			Status: models.ContentStatusDraft,
		}

		contentCreationResult := database.DB.Create(&newCampaignContent)
		if contentCreationResult.Error != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": contentCreationResult.Error.Error(),
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"content": newCampaignContent,
		})
	})

	// @Description Get a specific content
	// @Route api/protected/campaign/:id/content/:contentId
	contentGroup.Get("/:contentId", func(context *fiber.Ctx) error {
		contentId, err := strconv.Atoi(context.Params("contentId"))
		if err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		content := models.Content{}
		result := database.DB.First(&content, contentId)
		if result.Error != nil {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		
		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Content created successfully",
			"content": content,
		})
	})

	// @Description Update Content Status
	// @Route api/protected/campaign/:id/content/:contentId/
	contentGroup.Patch("/:contentId/status", func(context *fiber.Ctx) error {
		contentId, err := strconv.Atoi(context.Params("contentId"))
		if err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		content := models.Content{}
		result := database.DB.First(&content, contentId)
		if result.Error != nil {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}

		content.Status = models.ContentStatusPublished
		patchingResult := database.DB.Save(&content)
		if patchingResult.Error != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": patchingResult.Error.Error(),
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Content status updated successfully",
		})
	})

	// @Description Delete Content
	// @Route api/protected/campaign/:id/content/:contentId
	contentGroup.Delete("/:contentId", func(context *fiber.Ctx) error {
		contentId, err := strconv.Atoi(context.Params("contentId"))
		if err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		result :=database.DB.Delete(&models.Content{}, contentId)
		if result.Error != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Content deleted successfully",
		})
	})
}