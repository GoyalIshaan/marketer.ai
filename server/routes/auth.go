package routes

import (
	"log"
	"marketer-ai-backend/database"
	"marketer-ai-backend/models"
	"marketer-ai-backend/validation"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	LoginRequest
}

func jwtGenerator(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email": user.Email,
		"username": user.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func cookieGenerator(tokenString string, context *fiber.Ctx) error {
	context.Cookie(&fiber.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure: false,
		SameSite: "Lax",
	})

	return nil
}

func LoginRouter(context *fiber.Ctx) error {
	loginRequest := new(LoginRequest)
	if err := context.BodyParser(loginRequest); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if (loginRequest.Email == "" || loginRequest.Password == "") {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	var user models.User
	result := database.DB.Where("email = ?", loginRequest.Email).First(&user)

	if result.Error != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
    if err != nil {
        return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid email or password",
        })
    }

	tokenString, err := jwtGenerator(user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = cookieGenerator(tokenString, context)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
	})
}


func RegisterRouter(context *fiber.Ctx) error {
	registerRequest := new(RegisterRequest)
	
	if err := context.BodyParser(registerRequest); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}


	var user = models.User{
        Username: registerRequest.Username,
        Email:    registerRequest.Email,
        Password: registerRequest.Password,
    }

	// Input validation
	if !validation.IsValidUserRequest(user) {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user request",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
    if err != nil {
        return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to process password",
        })
    }

    user.Password = string(hashedPassword)

	result := database.DB.Create(&user)

	if result.Error != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	tokenString, err := jwtGenerator(user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = cookieGenerator(tokenString, context)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Register successful",
	})
}