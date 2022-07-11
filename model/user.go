package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username" gorm:"not null;unique;uniqueIndex"`
	Nickname    string `json:"nickname" gorm:"not null"`
	Email       string `json:"email" gorm:"not null;unique;uniqueIndex"`
	Password    string `json:"-" gorm:"not null"` // Password encrypted with HASH
	Description string `json:"description" gorm:"size:256"`
	Role        int    `json:"role" gorm:"not null;default:0"` // Role 0:user, 1:admin
}
