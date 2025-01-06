package aihandlers

import (
	"fmt"
	"marketer-ai-backend/models"
	"strings"
)

func PromptGenerator(contentData models.PromptContentResponse) string {
	var basePrompt strings.Builder
	basePrompt.WriteString("Based on the following campaign details, generate a title and content:\n\n")
	
	basePrompt.WriteString("Campaign Details:\n")
	basePrompt.WriteString("Content Type: ")
	basePrompt.WriteString(string(contentData.ContentType))
	basePrompt.WriteString("\n")
	basePrompt.WriteString("Campaign Title: ")
	basePrompt.WriteString(contentData.Campaign.Title)
	basePrompt.WriteString("\n")
	basePrompt.WriteString("Description: ")
	basePrompt.WriteString(contentData.Campaign.Description)
	basePrompt.WriteString("\n")
	basePrompt.WriteString("Budget: ")
	basePrompt.WriteString(fmt.Sprintf("%f", contentData.Campaign.Budget))
	basePrompt.WriteString("\n")
	basePrompt.WriteString("Platform: ")
	for _, platform := range contentData.Campaign.Platform {
		basePrompt.WriteString(string(platform))
		basePrompt.WriteString(", ")
	}
	basePrompt.WriteString("\n")
	basePrompt.WriteString("Hashtags: ")
	for _, hashtag := range contentData.Hashtags {
		basePrompt.WriteString(string(hashtag))
		basePrompt.WriteString(", ")
	}
	basePrompt.WriteString("\n\n")
	
	basePrompt.WriteString("Title (up to 10 words):\n")
	basePrompt.WriteString("Content (up to 100 words):\n")
	
	return basePrompt.String()
}