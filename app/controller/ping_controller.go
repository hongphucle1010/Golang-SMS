package controller

import (
	"sms/app/service"
	"sms/pkg/response"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type PingController struct{
	PingService service.IPingService	
}

// @Summary Health check
// @Description Ping to get pong
// @Tags ping
// @Produce  json
// @Success 200 {object} response.SuccessResponse[any]
// @Failure 404 {object} response.ErrorResponse
// @Router /ping/ [get]
func (c *PingController) Get() (mvc.Result, mvc.Err) {
	message, err := c.PingService.Pong()
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
			ContentType: response.JsonContentType,
			Object: response.ErrorResponse{Message: "Failed to ping", Details: err.Error()},
		}, nil
	}
	return mvc.Response{
		Code: iris.StatusOK,
		ContentType: response.JsonContentType,
		Object: response.SuccessResponse[any]{
			Message: "Successfully pinged",
			Data:    iris.Map{"message": message},
		},
	}, nil
}