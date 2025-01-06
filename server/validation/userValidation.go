package validation

import (
	"marketer-ai-backend/database"
	"marketer-ai-backend/models"
	"regexp"
)

func IsValidUserRequest(userRequest models.User) bool {
	return IsValidEmail(userRequest.Email) && IsValidPassword(userRequest.Password) && IsUsernameUnique(userRequest.Username)
}

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	result := emailRegex.MatchString(email) && IsEmailUnique(email)

	return result
}

func IsEmailUnique(email string) bool {
	var existingUser models.User
	result := database.DB.Where("email = ?", email).First(&existingUser)

	return result.Error != nil
}

func IsValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 20 {
		return false
	}
	
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[@$!%*?&]`).MatchString(password)
	
	return hasUpper && hasLower && hasNumber && hasSpecial
}

func IsUsernameUnique(username string) bool {
	var existingUser models.User
	result := database.DB.Where("username = ?", username).First(&existingUser)

	return result.Error != nil
}

func IsValidUserId(id uint) bool {
	var existingUser models.User
	result := database.DB.First(&existingUser, id)

	return result.Error == nil
}