package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"title" gorm:"size:128;not null"`
	Content string `json:"content" gorm:"not null"`
	UserID  uint   `json:"userId" gorm:"size:64"`
	User    User
}
