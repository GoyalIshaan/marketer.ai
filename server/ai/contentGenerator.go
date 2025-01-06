package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)
func GenerateContent(prompt string) (string, error) {
	requestPayload := OpenAIRequest{
		Model: "gpt-4o",
		Prompt: prompt,
		MaxTokens: 500,
		Temperature: 0.8,
	}

	requestBody, err := json.Marshal(requestPayload)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+OpenAIKey)

	//Send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("OpenAI API returned status: %d", response.StatusCode)
	}

	//Read the response
	var parsedResponse OpenAIResponse
	if err := json.NewDecoder(response.Body).Decode(&parsedResponse); err != nil {
		return "", err
	}

	if len(parsedResponse.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI API")
	}

	return parsedResponse.Choices[0].Text, nil
}
