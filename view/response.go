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
响应json数据
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
}

func main() {
	//创建一个路由
	Route := gin.Default()
	//绑定路由规则和路由函数，访问index的路由，交由对应的函数去处理
	Route.GET("/string", _string)
	Route.GET("/json", _json)
	//启动监听，gin会把web服务启动在8080端口
	Route.Run(":8080")
}
