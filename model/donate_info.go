package model

import "gorm.io/gorm"

type DonateInfo struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name"`
	Message     string `gorm:"unique" json:"message"`
	Email     string `gorm:"default:无留言;;index:donateInfo_message" json:"email"`
	Payment   string `gorm:"not null" json:"payment"`
	Url       string `gorm:"not null" json:"url"`
	CreatorId uint64 `gorm:"default:0;not null;index:donateInfo_creator" json:"creatorId"`
}
