package models

import (
	"time"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
