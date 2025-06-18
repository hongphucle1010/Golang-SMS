package main

import (
	"sms/app/config"
	"sms/app/middleware"
	"sms/app/routes"

	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func main() {
	app := iris.New()
	config.InitConfig()
	app.Use(middleware.Logger)
	routes.Register(app)
	app.Listen(":" + viper.GetString("server.port"))
}
