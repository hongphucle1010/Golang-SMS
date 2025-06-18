package routes

import (
	"github.com/kataras/iris/v12"
)

func Register(app *iris.Application) {
	api := app.Party("/api")
	RegisterPingRoutes(api)
	RegisterStudentRoutes(api)
	RegisterTestRoutes(api)
}
