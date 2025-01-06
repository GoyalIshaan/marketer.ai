package ai

// ChatMessage represents a single message in a conversation.
type ChatMessage struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

// ChatCompletionRequest is the payload for the /v1/chat/completions endpoint.
type ChatCompletionRequest struct {
    Model       string        `json:"model"`
    Messages    []ChatMessage `json:"messages"`
    MaxTokens   int           `json:"max_tokens,omitempty"`
    Temperature float32       `json:"temperature,omitempty"`
}

// ChatCompletionResponse is the response from /v1/chat/completions.
type ChatCompletionResponse struct {
    Choices []struct {
        Message struct {
            Role    string `json:"role"`
            Content string `json:"content"`
        } `json:"message"`
    } `json:"choices"`
}
