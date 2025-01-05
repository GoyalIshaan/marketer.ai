package routes

import (
	"time"

	"marketer-ai-backend/database"
	"marketer-ai-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var jwtSecret = []byte("secret")

func AuthRouter(context *fiber.Ctx) error {
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

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
	})
}
