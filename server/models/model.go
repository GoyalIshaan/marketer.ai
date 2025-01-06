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
	Contents []Content `json:"contents"`
}

type Content struct {
	ID uint `json:"id" gorm:"primaryKey"`
	CampaignID uint `json:"campaign_id" gorm:"not null"`
	Campaign Campaign `json:"campaign" gorm:"foreignKey:CampaignID"`
	Title string `json:"title"`
	Content string `json:"content"`
	ContentType ContentType `json:"content_type"`
	Status ContentStatus `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Hashtags pq.StringArray `json:"hashtags" gorm:"type:text[]"`
}