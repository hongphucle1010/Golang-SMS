package controller

import (
	"sms/app/service"
	"sms/pkg/response"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type TestController struct {
	TestService service.ITestService
}

// @Summary Get test
// @Description Get test
// @Tags test
// @Produce  json
// @Success 200 {object} response.SuccessResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /test/ [get]
func (c *TestController) Get() mvc.Result {
	test, err := c.TestService.GetTest()
	if err != nil {
		return mvc.Response{
			Code:        iris.StatusInternalServerError,
			ContentType: response.JsonContentType,
			Object:      response.ErrorResponse{Message: "Failed to fetch test", Details: err.Error()},
		}
	}
	return mvc.Response{
		Code:        iris.StatusOK,
		ContentType: response.JsonContentType,
		Object: response.SuccessResponse{
			Message: "Successfully fetched test",
			Data:    test,
		},
	}
}
