package ai

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var OpenAIKey string
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	OpenAIKey = os.Getenv("OPEN_AI_KEY")
}
