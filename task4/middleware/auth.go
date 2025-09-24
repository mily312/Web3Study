package middleware

import (
	"github.com/gin-gonic/gin"
)

// 鉴权中间件
func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// TODO
		c.Next()
	}
}
