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
	// Initialize config
	config.InitConfig()

	// Initialize Iris
	app := iris.New()
	app.Use(middleware.Recover)
	app.Use(middleware.Logger)

	// Register routes
	routes.Register(app)

	// Start server
	app.Listen(":" + viper.GetString("server.port"))
}
