package models

import (
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null;unique" json:"name" binding:"required" form:"name"`
	Blogs []Blog `gorm:"many2many:blog_topics;constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
}

type TopicResponse struct {
	Name  string         `json:"name"`
	Blogs []BlogResponse `json:"blogs"`
}
