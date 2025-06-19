package handlers

import "github.com/kataras/iris/v12"

type Student struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// GetStudents godoc
// @Summary Get students
// @Description Get details of a student
// @Tags students
// @Accept  json
// @Produce  json
// @Success 200 {object} Student
// @Failure 404 {object} ErrorResponse
// @Router /students/ [get]
func GetStudents(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": []Student{
		{ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
		{ID: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
		{ID: 3, Name: "Alice Johnson", Email: "alice.johnson@example.com"},
	}})
}	