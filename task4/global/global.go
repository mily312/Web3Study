package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger          *zap.SugaredLogger
	DB              *gorm.DB
	LOGIN_USER_ID   = "LOGIN_USER_ID"
	LOGIN_USER_NAME = "LOGIN_USER_NAME"
)
