package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	route := gin.Default()

	//请求头的获取方式
	route.GET("/", func(context *gin.Context) {

		//请求头大小写不区分，单词与单词之间用  —  连接，且返回切片中第一个数据
		fmt.Println(context.GetHeader("User-Agent"))
		fmt.Println(context.GetHeader("user-Agent"))

		//自定义请求头
		fmt.Println(context.Request.Header.Get("Token"))

		context.JSON(http.StatusOK, gin.H{"msg": "成功了"})
	})

	//爬虫和用户的区别对待
	route.GET("/index", func(context *gin.Context) {
		userAgent := context.GetHeader("User-Agent")
		//用正则法去匹配
		//用字符串去匹配
		if strings.Contains(userAgent, "python") {
			context.JSON(http.StatusOK, gin.H{"data": "这是响应给爬虫的数据！"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": "这是相应给用户的数据！"})
	})

	//设置响应头
	route.GET("/res", func(context *gin.Context) {
		context.Header("Token", "5r7t69yhjakfg89fdkdlfhibj")
		//context.Header("Content", "application/json; charset=utf-8")  改变请求头展示类型text\json
		//context.Header("Content", "application/text; charset=utf-8")
		context.JSON(http.StatusOK, gin.H{"data": "看看响应头"})
	})
	route.Run(":8080")
}
