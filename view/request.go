package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func _user(ctx *gin.Context) {
	fmt.Println(ctx.Query("user"))               //接受单个值用Query
	fmt.Println(ctx.GetQuery("user"))            //判断传没传，传了显示true，没传是false，传一个空值也是true注意
	fmt.Println(ctx.QueryArray("user"))          //接收多个值用QueryArray
	fmt.Println(ctx.DefaultQuery("addr", "大上海")) //接收一个默认值，没做此属性POST则按默认值传值，若是空字符串则是空字符串，另外传什么是什么
}

func _param(ctx *gin.Context) {
	fmt.Println(http.StatusOK, ctx.Param("user_id"))
	fmt.Println(http.StatusOK, ctx.Param("book_id"))
}

func _form(ctx *gin.Context) {
	name := ctx.PostForm("name")
	fmt.Println(http.StatusOK, name) // POST单个from属性
	names := ctx.PostFormArray("name")
	fmt.Println(http.StatusOK, names) // POST一个属性数组
	addr := ctx.DefaultPostForm("addr", "北京")
	fmt.Println(http.StatusOK, addr) // POST一个默认值，没做此属性POST则按默认值传值，若是空字符串则是空字符串，另外传什么是什么
	froms, err := ctx.MultipartForm()
	fmt.Println(http.StatusOK, froms, err) // 接受所有的from参数，包括文件
}

func _raw(ctx *gin.Context) {
	grd, _ := ctx.GetRawData()
	fmt.Println(string(grd))
}

func main() {
	route := gin.Default()

	route.GET("/user", _user)
	route.GET("/param/:user_id", _param)
	route.GET("/param/:user_id/:book_id", _param)
	route.POST("/form", _form)
	route.POST("/raw", _raw)

	route.Run(":8080")
}
