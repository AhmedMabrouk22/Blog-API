package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	BlogID  uint   `gorm:"not null"`
	UserID  uint   `grom:"not null"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CommentRes struct {
	ID      uint          `json:"id"`
	Content string        `json:"content"`
	Author  CommentAuthor `json:"author"`
}

type CommentAuthor struct {
	ID    uint   `json:"author_id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
