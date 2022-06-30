package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username" gorm:"not null;unique;index:user_name"`
	Nickname    string `json:"nickname" gorm:"not null"`
	Email       string `json:"email" gorm:"not null;unique;index:user_email"`
	Password    string `json:"-" gorm:"not null"` // Password encrypted with HASH
	Description string `json:"description"`
	Role        int    `json:"role" gorm:"not null;default:0"` // Role 0:user, 1:admin
}
