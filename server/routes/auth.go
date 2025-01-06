package routes

import (
	"marketer-ai-backend/database"
	"marketer-ai-backend/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	LoginRequest
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func tokenGenerator(user models.User, context *fiber.Ctx) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	context.Cookie(&fiber.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure: true,
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

	if user.Password != loginRequest.Password {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	err := tokenGenerator(user, context)

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
		Email: registerRequest.Email,
		Password: registerRequest.Password,
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	err := tokenGenerator(user, context)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Register successful",
	})
}