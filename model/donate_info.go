package model

import "gorm.io/gorm"

type DonateInfo struct {
	gorm.Model
	Name    string `gorm:"not null;index" json:"name"`
	Payment string `gorm:"not null" json:"payment"`
	Url     string `gorm:"not null" json:"url"`
	Comment string `gorm:"size:256" json:"comment"`
	Author  uint   `gorm:"default:0;not null;index" json:"author"`
}
