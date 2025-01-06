package database

import (
	"log"
	"marketer-ai-backend/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	DB = db
	log.Println("Connected to database")
	Migrate(models.User{}, models.Campaign{}, models.Content{})
}

func Migrate(models ...interface{}) {
	DB.AutoMigrate(models...)
}
