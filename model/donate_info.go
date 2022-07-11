package model

import "gorm.io/gorm"

// DonateInfo 捐赠信息
type DonateInfo struct {
	gorm.Model
	Name    string `gorm:"not null;index" json:"name"`             // Name 昵称
	Payment string `gorm:"not null" json:"payment"`                // Payment 支付方式
	Url     string `gorm:"not null" json:"url"`                    // Url 支付链接
	Comment string `gorm:"size:256" json:"comment"`                // Comment 留言
	Author  uint   `gorm:"default:0;not null;index" json:"author"` // Author 作者用户ID
}
