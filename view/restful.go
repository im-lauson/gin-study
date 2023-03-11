package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleMoudel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// 封装一个接收函数
func _bindJson(ctx *gin.Context, obj any) (err error) {
	boby, _ := ctx.GetRawData()
	ContentType := ctx.GetHeader("Content-Type")
	switch ContentType {
	case "application/json":
		err = json.Unmarshal(boby, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

// 文章列表
func _getList(ctx *gin.Context) {
	articleMoudel := []ArticleMoudel{
		{"Go语言入门", "这篇文章是Go语言入门"},
		{"Java语言入门", "这篇文章是Java语言入门"},
		{"Html语言入门", "这篇文章是Html语言入门"},
		{"Css语言入门", "这篇文章是Css语言入门"},
	}
	ctx.JSON(http.StatusOK, Response{200, articleMoudel, "成功"})
}

// 文章详情
func _getDetail(ctx *gin.Context) {
	//动态获取param中的id
	fmt.Println(ctx.Param("id"))
	article := ArticleMoudel{
		"Go语言入门", "这篇文章是Go语言入门",
	}

	ctx.JSON(http.StatusOK, Response{200, article, "成功"})
}

//创建文章

func _create(ctx *gin.Context) {
	//接收前端传来的json数据

	var article ArticleMoudel

	err := _bindJson(ctx, &article)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, Response{200, article, "请求成功"})
}

// 更新文章
func _update(ctx *gin.Context) {
	fmt.Println(ctx.Param("id"))

	var article ArticleMoudel
	err := _bindJson(ctx, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, Response{200, article, "修改成功"})
}

// 删除文章
func _delete(ctx *gin.Context) {
	fmt.Println(ctx.Param("id"))

	ctx.JSON(http.StatusOK, Response{200, map[string]string{}, "删除成功"})
}

func main() {

	route := gin.Default()
	route.GET("/articles", _getList)       //文章列表
	route.GET("/articles/:id", _getDetail) //文章详情
	route.POST("/articles", _create)       //添加文章
	route.PUT("/articles/:id", _update)    //编辑文章
	route.DELETE("/articles/:id", _delete) //删除列表

	route.Run(":8080")
}
