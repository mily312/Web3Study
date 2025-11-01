package api

import (
	"BlogSystem/global"
	"BlogSystem/utils"
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 公共属性
type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

type RequestBody struct {
	Ctx     *gin.Context
	Dto     any
	BindUri bool
	BindAll bool
}

// 请求参数封装
func (m *BaseApi) BuildRequest(requestBody RequestBody) *BaseApi {

	var errResult error

	// 绑定请求上下文
	m.Ctx = requestBody.Ctx

	// 绑定请求数据
	if requestBody.Dto != nil {
		if requestBody.BindAll || requestBody.BindUri {
			errResult = utils.AppendError(errResult, m.Ctx.ShouldBindUri(requestBody.Dto))
		}

		if requestBody.BindAll || !requestBody.BindUri {
			errResult = utils.AppendError(errResult, m.Ctx.ShouldBind(requestBody.Dto))
		}

		if errResult != nil {
			errResult = m.ParseValidateErrors(errResult, requestBody.Dto)
			m.AddError(errResult)
			m.Fail(ResponseBody{
				Msg: m.GetErrors().Error(),
			})
		}
	}

	return m
}

// 添加错误信息
func (m *BaseApi) AddError(errNew error) {
	if errNew != nil {
		m.Errors = utils.AppendError(m.Errors, errNew)
	}

}

// 获取错误信息
func (m *BaseApi) GetErrors() error {
	return m.Errors
}

// 解析错误信息
func (m *BaseApi) ParseValidateErrors(errs error, target any) error {
	var errResult error

	errValidation, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs
	}

	// 通过反射获取指针指向元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidation {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}

		if errMessage == "" {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}

		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}

	return errResult
}

// 公共响应方法
func (m *BaseApi) Fail(resp ResponseBody) {
	Fail(m.Ctx, resp)
}

func (m *BaseApi) OK(resp ResponseBody) {
	Ok(m.Ctx, resp)
}

func (m *BaseApi) ServerFail(resp ResponseBody) {
	ServerFail(m.Ctx, resp)
}
