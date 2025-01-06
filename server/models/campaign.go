package models

import (
	"time"

	"github.com/lib/pq"
)

type CampaignStatus string

const (
	CampaignStatusPending CampaignStatus = "pending"
	CampaignStatusRunning CampaignStatus = "running"
	CampaignStatusCompleted CampaignStatus = "completed"
	CampaignStatusFailed CampaignStatus = "failed"
)

type Platform string
const (
	PlatformFacebook Platform = "facebook"
	PlatformInstagram Platform = "instagram"
	PlatformTwitter Platform = "twitter"
	PlatformLinkedin Platform = "linkedin"
	PlatformYoutube Platform = "youtube"
)

type CampaignRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Budget float64 `json:"budget"`
	Platform pq.StringArray `json:"platform" gorm:"type:text[]"`
	Status CampaignStatus `json:"status"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
}

type StatusChangeRequest struct {
	Status CampaignStatus `json:"status"`
}