package api

import (
	"BlogSystem/global"
	"BlogSystem/model"
	"BlogSystem/service"
	"BlogSystem/service/dto"
	"strconv"

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

func (postApi PostApi) PostListPaginate(c *gin.Context) {
	var postPaginate dto.PostPaginateDto

	//封装请求参数
	errRequest := postApi.BuildRequest(RequestBody{
		Ctx: c,
		Dto: &postPaginate,
	}).GetErrors()

	if errRequest != nil {
		return
	}

	posts, total, err := postApi.postService.PostListPaginate(postPaginate)

	if err != nil {
		postApi.Fail(ResponseBody{
			Msg: "查询文章列表失败",
		})

		return
	}

	postApi.OK(ResponseBody{
		Data:  posts,
		Total: total,
	})
}

func (postApi PostApi) PostDetail(c *gin.Context) {
	var postId = c.Param("postId")
	postInt, _ := strconv.Atoi(postId)
	posts, err := postApi.postService.PostDetail(uint(postInt))

	if err != nil {
		postApi.Fail(ResponseBody{
			Msg: "查询文章详情失败",
		})

		return
	}

	postApi.Ctx = c
	postApi.OK(ResponseBody{
		Data: posts,
	})
}

func (postApi PostApi) UpdatePost(c *gin.Context) {
	var updatePost model.UpdatePost
	err := postApi.BuildRequest(RequestBody{
		Ctx:     c,
		Dto:     &updatePost,
		BindAll: true,
	}).GetErrors()
	if err != nil {
		return
	}

	errUpdate := postApi.postService.UpdatePost(updatePost)
	if errUpdate != nil {
		postApi.Fail(ResponseBody{
			Msg: "文章修改失败",
		})

		return
	}

	postApi.OK(ResponseBody{
		Msg:  "文章修改成功",
		Code: 200,
	})

}

func (postApi PostApi) DeletePost(c *gin.Context) {
	strPostId := c.Param("postId")
	intPostId, _ := strconv.Atoi(strPostId)

	err := postApi.postService.DelPost(uint(intPostId))
	if err != nil {
		postApi.Fail(ResponseBody{
			Msg: "文章删除失败",
		})

		return
	}

	postApi.Ctx = c
	postApi.OK(ResponseBody{
		Msg:  "文章删除成功",
		Code: 200,
	})
}
