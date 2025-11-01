package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:64;unique;not null"`
	Password string `json:"password" gorm:"size:128;not null"`
	Email    string `json:"email" gorm:"size:128;unique;not null"`
}
