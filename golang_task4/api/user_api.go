package api

import (
	"BlogSystem/model"
	"BlogSystem/service"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
	userService *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi:     NewBaseApi(),
		userService: service.NewUserService(),
	}
}

// userApi error全部存到这里，因为方法内func (userApi UserApi)传递的是非指针类型，值传递。每次请求互不影响
var apiErrors error

func (userApi UserApi) Login(c *gin.Context) {
	var userParam model.User
	userApi.BuildRequest(RequestBody{
		Ctx: c,
		Dto: &userParam,
	})
	token, err := userApi.userService.Login(&userParam)
	if err != nil {
		userApi.Fail(ResponseBody{
			Msg: "login failed!" + err.Error(),
		})

		return
	}

	userApi.OK(ResponseBody{
		Msg:  "Login Success",
		Data: token,
	})
}

// 用户注册
func (userApi UserApi) Register(c *gin.Context) {
	var user model.User
	reqErr := userApi.BuildRequest(RequestBody{
		Ctx: c,
		Dto: &user,
	}).Errors

	if reqErr != nil {
		//错误记录到日志中
		userApi.Logger.Error(reqErr.Error())
		userApi.Fail(ResponseBody{
			Msg: reqErr.Error(),
		})

		return
	}

	registerErr := userApi.userService.Register(&user)
	if registerErr != nil {
		userApi.Logger.Error(registerErr.Error())
		userApi.Fail(ResponseBody{
			Msg: registerErr.Error(),
		})

		return
	}

	userApi.OK(ResponseBody{
		Msg: "Register success!",
	})
}
