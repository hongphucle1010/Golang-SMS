package routes

import (
	"github.com/kataras/iris/v12"
	"sms/app/handlers"
)

func RegisterTestRoutes(app iris.Party) {
	app.Get("/test", handlers.GetTest)
}