package routes

import (
	"github.com/kataras/iris/v12"
	"sms/app/handlers"
)

func RegisterPingRoutes(app iris.Party) {
	app.Get("/ping", handlers.Ping)
}
