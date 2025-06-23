package controller

import (
	"sms/app/model"
	"sms/app/service"
	"sms/pkg/response"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type StudentController struct {
	Ctx iris.Context
	StudentService service.IStudentService
}

// @Summary Get all students
// @Description Get all students
// @Tags students
// @Produce  json
// @Success 200 {object} response.SuccessResponse[[]model.Student]
// @Failure 500 {object} response.ErrorResponse
// @Router /students/ [get]
func (c *StudentController) Get() mvc.Result {
	students, err := c.StudentService.GetAllStudents()
	if err != nil {
        return response.ErrorResponse{
            Code:    iris.StatusInternalServerError,
            Message: "Failed to fetch students",
            Details: err.Error(),
        }
	}

	return response.SuccessResponse[[]model.Student]{
		Code:    iris.StatusOK,
		Message: "Successfully fetched students",
		Data:    students,
	}
}
