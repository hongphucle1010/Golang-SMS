package controller

import (
	"sms/pkg/response"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type PingController struct{}

func (c *PingController) Get() (mvc.Result, mvc.Err) {
	return mvc.Response{
		Code: iris.StatusOK,
		ContentType: response.JsonContentType,
		Object: response.SuccessResponse{
			Message: "pong",
			Data:    iris.Map{"message": "pong"},
		},
	}, nil
}