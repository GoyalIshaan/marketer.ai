package validation

import (
	"marketer-ai-backend/database"
	"marketer-ai-backend/models"
)

func IsValidCampaignRequest(campaignRequest models.CampaignRequest) bool {
	return IsValidTitle(campaignRequest.Title) && IsValidDescription(campaignRequest.Description) && IsValidBudget(campaignRequest.Budget) && IsValidPlatform(campaignRequest.Platform)
}

func IsValidCampaignId(id uint) bool {
	result := database.DB.First(&models.Campaign{}, id)
	return result.Error == nil
}

func IsValidTitle(title string) bool {
	return len(title) > 0
}

func IsValidDescription(description string) bool {
	return len(description) > 0
}

func IsValidBudget(budget float64) bool {
	return budget > 0
}

func IsValidPlatform(platforms []string) bool {
	 validPlatforms := map[string]bool{
        string(models.PlatformFacebook):   true,
        string(models.PlatformInstagram):  true,
        string(models.PlatformTwitter):    true,
        string(models.PlatformLinkedin):   true,
        string(models.PlatformYoutube):    true,
    }

	for _, platform := range platforms {
		if !validPlatforms[platform] {
			return false
		}
	}
	return true
}

func IsValidStatus(status models.CampaignStatus) bool {
	return status == models.CampaignStatusPending || status == models.CampaignStatusRunning || status == models.CampaignStatusCompleted || status == models.CampaignStatusFailed
}
