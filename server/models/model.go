package models

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Campaigns []Campaign `json:"campaigns"`
	CreatedAt time.Time `json:"created_at"`
}

type Campaign struct {
	ID uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"not null"`
	User User `json:"user" gorm:"foreignKey:UserID"`
	Title string `json:"title"`
	Description string `json:"description"`
	Budget float64 `json:"budget"`
	Platform pq.StringArray `json:"platform" gorm:"type:text[]"`
	Status CampaignStatus `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
}

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
