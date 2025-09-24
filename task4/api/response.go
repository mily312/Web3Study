package api

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
	Total  int64
}

// 自定义状态码返回
func getStatus(responseBody ResponseBody, defaultStatus int) int {
	if responseBody.Status == 0 {
		return defaultStatus
	} else {
		return responseBody.Status
	}
}

func Ok(ctx *gin.Context, responseBody ResponseBody) {
	HttpResponse(ctx, responseBody, getStatus(responseBody, http.StatusOK))
}

func Fail(ctx *gin.Context, responseBody ResponseBody) {
	HttpResponse(ctx, responseBody, getStatus(responseBody, http.StatusBadRequest))
}

// 通用相应函数
func HttpResponse(ctx *gin.Context, responseBody ResponseBody, status int) {
	if responseBody.isEmpty() {
		ctx.AbortWithStatus(status)
		return
	}

	ctx.AbortWithStatusJSON(status, responseBody)
}

// 判断结构体对象是不是空
func (res ResponseBody) isEmpty() bool {
	return reflect.DeepEqual(res, ResponseBody{})
}

func ServerFail(ctx *gin.Context, resp ResponseBody) {
	HttpResponse(ctx, resp, getStatus(resp, http.StatusInternalServerError))
}
