package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	//相应字符串 string
	context.String(200, "hello")
	context.String(http.StatusOK, "hello")
}

func main() {
	//创建一个路由
	route := gin.Default()

	//绑定路由规则和路由函数，访问index的路由，交由对应的函数去处理
	route.GET("/index", Index) //引用函数也可
	//route.GET("/index",func(context *gin.Context) {
	//	context.String(200, "hello")
	//})

	//启动监听，gin会把web服务启动在8080端口
	//route.Run(":8080")
	http.ListenAndServe(":8080", route)

}
