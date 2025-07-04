package routes

import (
	"sms/app/controller"
	"sms/app/repository"
	"sms/app/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Register(app *iris.Application) {
	PreflightHandler(app)				// Register Preflight Handler
	RegisterSwagger(app)				// Register Swagger

	api := app.Party("/api")

	mvc.New(api.Party("/ping")).Register(service.NewPingService).Handle(new(controller.PingController))
	mvc.New(api.Party("/students")).Register(
		repository.NewJsonStudentRepository,
		service.NewStudentService,
	).Handle(new(controller.StudentController))
	mvc.New(api.Party("/test")).Register(&service.TestService{}).Handle(new(controller.TestController))
}
