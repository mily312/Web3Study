package middleware

import (
	"BlogSystem/api"
	"BlogSystem/global"
	"BlogSystem/model"
	"BlogSystem/utils"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_INVALID_TOKEN_INVALID     = 10401
	ERR_CODE_INVALID_TOKEN_PARSE_ERROR = 10402
	ERR_CODE_INVALID_TOKEN_NOT_MATCHED = 10403
	ERR_CODE_INVALID_TOKEN_EXPIRED     = 10404
	ERR_CODE_INVALID_TOKEN_RENEW_ERROR = 10405
	TOKEN_NAME                         = "Authorization"
	TOKEN_PREFIX                       = "Bearer: "
	RENEW_TOKEN_DURATION               = 10 * 60 * time.Second
)

func tokenErr(c *gin.Context, code int) {
	api.Fail(c, api.ResponseBody{
		Status: http.StatusUnauthorized,
		Code:   code,
		Msg:    "Invalid Token",
	})
}

// 鉴权中间件
func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader(TOKEN_NAME)

		// Token不存在, 直接返回
		if token == "" || !strings.HasPrefix(token, TOKEN_PREFIX) {
			tokenErr(c, ERR_CODE_INVALID_TOKEN_INVALID)
			return
		}

		// Token无法解析, 直接返回
		token = token[len(TOKEN_PREFIX):]
		iJwtCustClaims, err := utils.ParseToken(token)
		nUserId := iJwtCustClaims.ID
		if err != nil || nUserId == 0 {
			fmt.Println(err.Error())
			tokenErr(c, ERR_CODE_INVALID_TOKEN_PARSE_ERROR)
			return
		}

		// 查询数据库，根据id和name查询是否有该记录
		var user model.User
		var totalNum int64
		recordErr := global.DB.Model(&model.User{}).
			Where("id = ? and username = ?", nUserId, iJwtCustClaims.Name).
			First(&user).Count(&totalNum).Error
		if recordErr != nil || totalNum == 0 {
			tokenErr(c, ERR_CODE_INVALID_TOKEN_NOT_MATCHED)
			return
		}

		// 将用户信息存入上下文, 方便后续处理继续使用
		c.Set(global.LOGIN_USER_ID, nUserId)
		c.Set(global.LOGIN_USER_NAME, iJwtCustClaims.Name)

		c.Next()
	}
}
