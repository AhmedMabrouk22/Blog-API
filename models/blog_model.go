package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title      string    `gorm:"type:varchar;not null" json:"title" binding:"required" form:"title"`
	Content    string    `gorm:"type:text; not null" json:"content" binding:"required" form:"content"`
	ImageCover string    `gorm:"type:varchar; not null"`
	AuthorID   uint      `gorm:"not null"`
	Author     User      `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Topics     []Topic   `gorm:"many2many:blog_topics;constraint:OnUpdate:SET NULL,OnDelete:CASCADE;"`
	Comments   []Comment `gorm:"foreignKey:BlogID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Likes      []Like    `gorm:"foreignKey:BlogID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type BlogRequest struct {
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
	ImageCover string
	AuthorID   uint
	Topics     []string `json:"Topics" form:"topics"`
}

type BlogResponse struct {
	ID         uint          `json:"id"`
	Title      string        `json:"title"`
	Content    string        `json:"content"`
	ImageCover string        `json:"imageCover"`
	Author     BlogAuthorRes `json:"author"`
	Topics     []Topic       `json:"topics"`
	Comments   []CommentRes  `json:"comments"`
	Likes      []LikeRes     `json:"likes"`
}

type BlogAuthorRes struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
