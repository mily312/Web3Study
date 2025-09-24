package api

import (
	"BlogSystem/global"
	"BlogSystem/model"
	"BlogSystem/service"

	"github.com/gin-gonic/gin"
)

type PostApi struct {
	BaseApi
	postService *service.PostService
}

func NewPostApi() PostApi {
	return PostApi{
		BaseApi:     NewBaseApi(),
		postService: service.NewPostService(),
	}
}

func (postApi PostApi) AddPost(c *gin.Context) {
	var post model.Post
	postApi.BuildRequest(RequestBody{
		Ctx: c,
		Dto: &post,
	})

	//userId取当前登录用户
	userId := c.GetUint(global.LOGIN_USER_ID)
	post.UserID = uint(userId)

	err := postApi.postService.AddPost(&post)

	if err != nil {
		postApi.Fail(ResponseBody{
			Msg:  "add post failed",
			Data: err.Error(),
		})

		return
	}

	postApi.OK(ResponseBody{
		Msg:  "add post success!",
		Data: post,
	})
}
