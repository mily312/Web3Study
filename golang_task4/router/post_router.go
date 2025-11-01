package router

import (
	"BlogSystem/api"
	"BlogSystem/middleware"

	"github.com/gin-gonic/gin"
)

func InitPostRouter() {
	//往切面里面添加元素，元素为：IFnRegistRouter,相当于调用RegistRouter方法
	RegistRouter(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		postApi := api.NewPostApi()

		rgPublicGroup := rgPublic.Group("post") //文章接口有的需要鉴权，有的不需要
		{
			rgPublicGroup.POST("pageList", postApi.PostListPaginate)
			rgPublicGroup.POST("postDetail/:postId", postApi.PostDetail)
		}

		rgAuthGroup := rgAuth.Group("post")
		{
			rgAuthGroup.POST("add", postApi.AddPost) //文章新增
		}

		//修改和删除只有文章作者有权限
		rgAuthGroup.Use(middleware.OnlyOwner())
		{
			rgAuthGroup.POST("update/:postId", postApi.UpdatePost) //文章修改
			rgAuthGroup.POST("delete/:postId", postApi.DeletePost) //文章删除
		}

	})
}
