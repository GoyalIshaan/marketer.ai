package ai

// OpenAIRequest payload
type OpenAIRequest struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
	MaxTokens int `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
}

// OpenAIResponse payload
type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}