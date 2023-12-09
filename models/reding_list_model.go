package models

import "gorm.io/gorm"

type ReadingList struct {
	gorm.Model
	Name   string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
	Blogs  []Blog `gorm:"many2many:user_reading_lists;"`
}

type ReadingListRes struct {
	Name        string `json:"name"`
	BlogsNumber int    `json:"blogs_number"`
}

type BlogReadingListRes struct {
	BlogID    uint   `json:"blog_id"`
	BlogName  string `json:"blog_name"`
	BlogImage string `json:"blog_image"`
}
