package middleware

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwt_secret []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	jwt_secret = []byte(os.Getenv("JWT_SECRET"))
}

func Protected() fiber.Handler {
	return func(context *fiber.Ctx) error {
		token := context.Cookies("token")
		if token == "" {
			return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Printf("Unexpected signing method: %v", token.Method.Alg())
				return nil, fiber.NewError(fiber.StatusUnauthorized, "unexpected signing method")
			}
			return jwt_secret, nil
		})

		if err != nil {
			log.Printf("Token parsing error details: %v", err)
			return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
				"details": err.Error(),
			})
		}

		if !parsedToken.Valid {
			log.Printf("Token validation failed")
			return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}


		claims := parsedToken.Claims.(jwt.MapClaims)

		userId := uint(claims["user_id"].(float64))
		email := claims["email"].(string)
		username := claims["username"].(string)
		
		context.Locals("user_id", userId)
		context.Locals("username", username)
		context.Locals("email", email)

		return context.Next()
	}
}
