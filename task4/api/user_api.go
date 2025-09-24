package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

func (userApi UserApi) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
