package controller

import (
	"sms/app/model"
	"sms/pkg/response"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type StudentController struct{}

// @Summary Get all students
// @Description Get all students
// @Tags students
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Student
// @Failure 404 {object} response.ErrorResponse
// @Router /students/ [get]
func (c *StudentController) Get() mvc.Result {
	successResponse := response.SuccessResponse{
		Message: "Successfully fetched students",
        Data: []model.Student{
            {
                ID: 1,
                Name: "John Doe",
                Email: "john.doe@example.com",
            },
            {
                ID: 2,
                Name: "Jane Doe",
                Email: "jane.doe@example.com",
            },
        },
	}

	return mvc.Response{
        Code: iris.StatusOK,
        ContentType: response.JsonContentType,
		Object: successResponse,
	}
}
