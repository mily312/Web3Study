package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"title" gorm:"size:128;not null" binding:"required"`
	Content string `json:"content" gorm:"not null" binding:"required"`
	UserID  uint   `json:"userId" gorm:"size:64;not null"`
	User    User
}
