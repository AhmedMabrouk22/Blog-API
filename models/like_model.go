package models

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BlogID uint `gorm:"not null"`
}

type LikeRes struct {
	ID   uint     `json:"id"`
	User LikeUser `json:"user"`
}

type LikeUser struct {
	ID    uint   `json:"author_id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
