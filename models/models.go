package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string
	Tweets   []Tweet `json:"tweets"`
}

type Tweet struct {
	gorm.Model
	UserId  int
	Content string
	User    User
}
