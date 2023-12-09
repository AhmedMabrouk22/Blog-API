package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required" gorm:"type:varchar(250);not null"`
	Email    string `json:"email" binding:"required" gorm:"type:varchar(250);not null;unique"`
	Password string `json:"password" binding:"required,min=8" gorm:"not null"`
	Image    string
	Blogs    []Blog `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserResponse struct {
	Name  string         `json:"name"`
	Image string         `json:"image"`
	Blogs []BlogResponse `json:"blogs"`
}
