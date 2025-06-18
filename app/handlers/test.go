package handlers

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func GetTest(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": viper.GetString("TEST_ENV")})
}