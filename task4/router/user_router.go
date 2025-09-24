package router

import (
	"BlogSystem/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter() {
	//往切面里面添加元素，元素为：IFnRegistRouter,相当于调用RegistRouter方法
	RegistRouter(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()

		rgPublic.POST("login", userApi.Login)

		// rgAuthUser := rgAuth.Group("/user")
		// {
		// 	rgAuthUser.POST("/addUser", userApi.AddUser)
		// 	rgAuthUser.GET("/queryUserById", userApi.QueryUserById)
		// 	rgAuthUser.POST("/list", userApi.QueryUserList)
		// 	rgAuthUser.PUT("/:id", userApi.UpdateUser)
		// }

	})
}
