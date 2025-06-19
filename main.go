package main

import (
	"sms/app/config"
	"sms/app/middleware"
	"sms/app/routes"
	_ "sms/docs"

	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

// @title Swagger Student Management System API
// @version 1.0
// @description This is a sample server Student Management System.
// @termsOfService http://swagger.io/terms/

// @contact.name HCMUT Team at Fsoft
// @contact.email hongphucle1010@gmail.com

// @host localhost:8080
// @BasePath /api
func main() {
	app := iris.New()
	config.InitConfig()
	app.Use(middleware.Logger)

	routes.Register(app)
	app.Listen(":" + viper.GetString("server.port"))
}
