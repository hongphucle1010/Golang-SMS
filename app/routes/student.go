package routes

import (
	"github.com/kataras/iris/v12"
	"sms/app/handlers"
)

func RegisterStudentRoutes(app iris.Party) {
	app.Get("/students", handlers.GetStudents)
}