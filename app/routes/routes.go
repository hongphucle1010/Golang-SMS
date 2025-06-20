package routes

import (
	"sms/app/controller"
	"sms/app/repository"
	"sms/app/service"

	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Register(app *iris.Application) {
	app.Get("/swagger/{any:path}", func(ctx iris.Context) {
		swagger.CustomWrapHandler(&swagger.Config{
			URL: "http://" + ctx.Host() + "/swagger/doc.json",
		}, swaggerFiles.Handler)(ctx)
	})

	api := app.Party("/api")

	mvc.New(api.Party("/ping")).Register(service.NewPingService).Handle(new(controller.PingController))
	mvc.New(api.Party("/students")).Register(
		repository.NewJsonStudentRepository,
		service.NewStudentService,
	).Handle(new(controller.StudentController))
	mvc.New(api.Party("/test")).Register(&service.TestService{}).Handle(new(controller.TestController))
}
