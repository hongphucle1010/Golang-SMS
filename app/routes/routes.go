package routes

import (
	"sms/app/controller"

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

	mvc.New(api.Party("/ping")).Handle(new(controller.PingController))
	mvc.New(api.Party("/students")).Handle(new(controller.StudentController))
	mvc.New(api.Party("/test")).Handle(new(controller.TestController))
}