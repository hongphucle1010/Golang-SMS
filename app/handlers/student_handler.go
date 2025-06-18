package handlers

import "github.com/kataras/iris/v12"

func GetStudents(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "students"})
}	