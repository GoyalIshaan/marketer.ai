package models

import "github.com/lib/pq"

type ContentStatus string

const (
	ContentStatusDraft ContentStatus = "draft"
	ContentStatusScheduled ContentStatus = "scheduled"
	ContentStatusPublished ContentStatus = "published"
	ContentStatusArchived ContentStatus = "archived"
)

type ContentType string

const (
	ContentTypeAdCopy ContentType = "ad_copy"
	ContentTypeBlog ContentType = "blog"
	ContentTypePoem ContentType = "poem"
	ContentTypeSocialMediaPost ContentType = "social_media_post"
	ContentTypeHashtag ContentType = "hashtag"
)

type GenerateContentRequest struct {
	CampaignID uint `json:"campaign_id"`
	ContentType ContentType `json:"content_type"`
	Hashtags pq.StringArray `json:"hashtags" gorm:"type:text[]"`
}

type PromptContentResponse struct {
	ContentType ContentType `json:"content_type"`
	Campaign Campaign `json:"campaign"`
	Hashtags pq.StringArray `json:"hashtags" gorm:"type:text[]"`
}