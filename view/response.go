package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
响应字符串 string
*/
func _string(c *gin.Context) {
	c.String(http.StatusOK, " 你好啊！")
}

/*
*
响应json结构体是最常用的（重点）
*/
func _json(c *gin.Context) {
	//定义一个结构体
	type UserInfo struct {
		//属性包括username、age、password
		Username string `json:"user_name"`
		Age      int    `json:"age"`
		Password string `json:"password"` //  “-”代表不去渲染数据，不去进行json序列化
	}
	//实例化user信息
	user := UserInfo{"songsong", 23, "123456"}
	c.JSON(200, user)

	//json响应map
	//userMap := map[string]string{
	//	"user_name": "liusong",
	//	"age":       "73",
	//}
	//c.JSON(200, userMap)

	//直接响应json
	//c.JSON(http.StatusOK, gin.H{"username": "liusong", "age": "345"})
}

/*
*
响应xml
*/
func _xml(x *gin.Context) {
	type UserInfo struct {
		//属性包括username、age、password
		Username string `json:"user_name"`
		Age      int    `json:"age"`
		Password string `json:"password"` //  “-”代表不去渲染数据，不去进行json序列化
	}
	user := UserInfo{"songsong", 23, "123456"}

	x.XML(http.StatusOK, user)
}

/*
*
响应xml、html、yaml
*/
func _redirect(c *gin.Context) {
	c.Redirect(302, "http://www.baidu.com")
}
func Index(context *gin.Context) {
	//相应字符串 string
	context.String(200, "hello")
	context.String(http.StatusOK, "hello")
}
func main() {
	//创建一个路由
	Route := gin.Default()
	//绑定路由规则和路由函数，访问index的路由，交由对应的函数去处理
	Route.GET("/string", _string)
	Route.GET("/json", _json)
	Route.GET("/xml", _xml)
	//重定向响应
	Route.GET("/baidu", _redirect)
	//绑定路由规则和路由函数，访问file路径，交由相应的函数去处理
	//在golang中没有相对路径，只有相对项目路径
	Route.StaticFile("/logo", "./static/logo.jpg")

	//配置单个文件
	Route.StaticFS("/static", http.Dir("static/static"))

	//绑定路由规则和路由函数，访问index的路由，交由对应的函数去处理
	Route.GET("/index", Index) //引用函数也可
	//route.GET("/index",func(context *gin.Context) {
	//	context.String(200, "hello")
	//})

	//启动监听，gin会把web服务启动在8080端口
	//route.Run(":8080")
	//http.ListenAndServe(":8080", Route)

	//启动监听，gin会把web服务启动在8080端口
	Route.Run(":8080")
}
