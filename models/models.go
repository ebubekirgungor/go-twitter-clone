package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `gorm:"uniqueIndex;not null" json:"username"`
	Email    string  `gorm:"uniqueIndex;not null" json:"email"`
	Password string  `gorm:"not null"`
	Name     string  `json:"name"`
	Tweets   []Tweet `json:"tweets"`
}

type Tweet struct {
	gorm.Model
	UserId  int    `gorm:"not null" json:"user_id"`
	Content string `gorm:"not null" json:"content"`
	User    User
}
