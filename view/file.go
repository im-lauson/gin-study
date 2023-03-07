package main

import "github.com/gin-gonic/gin"

//
//func _file(ctx gin.Context)  {
//	c.
//}

func main() {

	//创建一个路由
	Route := gin.Default()

	//绑定路由规则和路由函数，访问file路径，交由相应的函数去处理
	Route.StaticFile("/static/logo", "./static/logo.jpg")

	//Route.GET("file",_filr)
	//启动监听，Gin把服务启动在8080端口
	Route.Run(":8080")
}
