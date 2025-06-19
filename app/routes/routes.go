package routes

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
)

func Register(app *iris.Application) {
	app.Get("/swagger/{any:path}", func(ctx iris.Context) {
		swagger.CustomWrapHandler(&swagger.Config{
			URL: "http://" + ctx.Host() + "/swagger/doc.json",
		}, swaggerFiles.Handler)(ctx)
	})
	
	api := app.Party("/api")
	RegisterPingRoutes(api)
	RegisterStudentRoutes(api)
	RegisterTestRoutes(api)
}
