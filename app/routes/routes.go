package routes

import (
	"sms/app/controller"
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

	mvc.New(api.Party("/ping")).Handle(new(controller.PingController)).Register(&service.PingService{})
	mvc.New(api.Party("/students")).Handle(new(controller.StudentController)).Register(&service.StudentService{})
	mvc.New(api.Party("/test")).Handle(new(controller.TestController)).Register(&service.TestService{})
}
