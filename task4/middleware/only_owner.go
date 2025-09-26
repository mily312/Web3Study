package middleware

import (
	"BlogSystem/global"
	"BlogSystem/model"
	"BlogSystem/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	NO_FOUND_POST = 20001
	NO_PERMISSION = 20002
)

// 只有文章作者可以操作
func OnlyOwner() func(c *gin.Context) {
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

		//查询数据库文章作者id和token解析的id是否一致
		var postId = c.Param("postId")
		postInt, _ := strconv.Atoi(postId)

		var postInfo model.Post
		errQuery := global.DB.Model(&model.Post{}).First(&postInfo, postInt).Error
		if errQuery != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Msg":  "未查找到文章信息",
				"Code": NO_FOUND_POST,
			})
			return
		}

		if nUserId != postInfo.UserID {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Msg":  "当前用户无权限",
				"Code": NO_PERMISSION,
			})
			return
		}

		c.Next()
	}
}
