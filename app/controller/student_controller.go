package controller

import (
	"sms/app/model"
	"sms/app/service"
	"sms/pkg/response"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type StudentController struct {
	StudentService service.IStudentService
}

// @Summary Get all students
// @Description Get all students
// @Tags students
// @Produce  json
// @Success 200 {object} response.SuccessResponse[[]model.Student]
// @Failure 404 {object} response.ErrorResponse
// @Router /students/ [get]
func (c *StudentController) Get() mvc.Result {
	students, err := c.StudentService.GetAllStudents()
	if err != nil {
		return mvc.Response{
			Code:        iris.StatusInternalServerError,
			ContentType: response.JsonContentType,
			Object:      response.ErrorResponse{Message: "Failed to fetch students", Details: err.Error()},
		}
	}

	successResponse := response.SuccessResponse[[]model.Student]{
		Message: "Successfully fetched students",
		Data:    students,
	}

	return mvc.Response{
		Code:        iris.StatusOK,
		ContentType: response.JsonContentType,
		Object:      successResponse,
	}
}
