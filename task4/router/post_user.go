package router

import (
	"BlogSystem/api"

	"github.com/gin-gonic/gin"
)

func InitPostRouter() {
	//往切面里面添加元素，元素为：IFnRegistRouter,相当于调用RegistRouter方法
	RegistRouter(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		postApi := api.NewPostApi()

		rgAuthGroup := rgAuth.Group("post")
		{
			rgAuthGroup.POST("add", postApi.AddPost) //文章新增
		}

	})
}
