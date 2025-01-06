package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GenerateContent(prompt string) (string, error) {
    requestPayload := ChatCompletionRequest{
        Model: "gpt-3.5-turbo",
        Messages: []ChatMessage{
            {Role: "user", Content: prompt},
        },
        MaxTokens:   500,
        Temperature: 0.8,
    }

    requestBody, err := json.Marshal(requestPayload)
	if err != nil {
		log.Println("Error marshaling request payload:", err)
		return "", err
	}

	log.Println("Request Body:", string(requestBody))

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("Error creating new request:", err)
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+OpenAIKey)

	// Send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error sending request:", err)
		return "", err
	}
	defer response.Body.Close()


	if response.StatusCode != http.StatusOK {
		log.Println("OpenAI API returned status:", response.StatusCode)
		return "", fmt.Errorf("OpenAI API returned status: %d", response.StatusCode)
	}

	// Read the response
	var parsedResponse ChatCompletionResponse
	if err := json.NewDecoder(response.Body).Decode(&parsedResponse); err != nil {
		log.Println("Error decoding response:", err)
		return "", err
	}

	if len(parsedResponse.Choices) == 0 {
		log.Println("No response from OpenAI API")
		return "", fmt.Errorf("no response from OpenAI API")
	}

	return parsedResponse.Choices[0].Message.Content, nil
}
