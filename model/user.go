package model

import "gorm.io/gorm"

// User 用户信息
type User struct {
	gorm.Model
	Username    string `json:"username" gorm:"not null;unique;uniqueIndex"` // Username 用户名 不可变/登陆凭据/唯一
	Nickname    string `json:"nickname" gorm:"not null"`                    // Nickname 昵称 可变
	Email       string `json:"email" gorm:"not null;unique;uniqueIndex"`    // Email 用户邮箱
	Password    string `json:"-" gorm:"not null"`                           // Password 哈希加密
	Description string `json:"description" gorm:"size:256"`                 // Description 用户简介
	Role        int    `json:"role" gorm:"not null;default:0"`              // Role 0:user, 1:admin
}
