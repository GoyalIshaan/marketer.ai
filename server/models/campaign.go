package models

import (
	"time"

	"github.com/lib/pq"
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