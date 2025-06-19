package controller

import (
	"sms/pkg/response"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type TestController struct{}

func (c *TestController) Get() mvc.Result {
	return mvc.Response{
		Code: iris.StatusOK,
		ContentType: response.JsonContentType,
		Object: response.SuccessResponse{
			Message: "pong",
			Data:    iris.Map{"message": "pong"},
		},
	}
}