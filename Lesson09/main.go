package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 加载静态文件
	r.Static("/xxx", "./statics")

	// gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	//Gin框架中使用LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染。
	r.LoadHTMLFiles("templates/posts/index.tmpl", "templates/users/index.tmpl")
	//r.LoadHTMLGlob("templates/**/*")

	// GET：请求方式；/index：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "user/index",
		})
	})

	r.Run()
}
