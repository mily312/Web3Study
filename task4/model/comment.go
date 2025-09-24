package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint   `json:"userId" gorm:"size:64"`
	User    User
	PostID  uint `json:"postId" gorm:"size:64"`
	Post    Post
}
