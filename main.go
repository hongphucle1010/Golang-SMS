package main

import (
	"sms/app/config"
	"sms/app/middleware"
	"sms/app/routes"
	_ "sms/docs"

	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

// @title Swagger UI Hệ thống quản lý sinh viên
// @version 1.0
// @description Tài liệu API của hệ thống quản lý sinh viên
// @termsOfService http://swagger.io/terms/

// @contact.name HCMUT Team tại FPT Software HCM
// @contact.email hongphucle1010@gmail.com

// @BasePath /api
func main() {
	// Initialize config
	config.InitConfig()

	// Initialize Iris
	app := iris.New()
	app.Use(middleware.Recover)
	app.Use(middleware.Logger)
	app.Use(middleware.Cors())

	// Register routes
	routes.Register(app)

	// Start server
	app.Listen(":" + viper.GetString("server.port"))
}
