package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 声明一个结构体
type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func main() {
	route := gin.Default()

	route.POST("/bind", func(context *gin.Context) {

		//实例一个用户
		var userinfo UserInfo

		err := context.ShouldBindJSON(&userinfo)
		//自带一个类型判断
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"msg": "你错了！"})
			return
		}
		context.JSON(200, userinfo)
	})

	route.POST("/query", func(context *gin.Context) {

		//实例一个用户
		var userinfo UserInfo

		err := context.ShouldBindQuery(&userinfo)
		//自带一个类型判断
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"msg": "你错了！"})
			return
		}
		context.JSON(200, userinfo)
	})

	route.POST("/uri/:name/:age/:sex", func(context *gin.Context) {

		//实例一个用户
		var userinfo UserInfo

		err := context.ShouldBindUri(&userinfo)
		//自带一个类型判断
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"msg": "你错了！"})
			return
		}
		context.JSON(200, userinfo)
	})

	route.POST("/from", func(context *gin.Context) {

		//实例一个用户
		var userinfo UserInfo

		err := context.ShouldBind(&userinfo)
		//自带一个类型判断
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"msg": "你错了！"})
			return
		}
		context.JSON(200, userinfo)
	})

	route.Run(":8080")
}
