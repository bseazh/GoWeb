package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		// 跳转到sogo.com
		c.Redirect(http.StatusMovedPermanently, "http://sogo.com")
	})

	r.GET("/a", func(c *gin.Context) {
		// 跳转/b对应的路由处理函数
		c.Request.URL.Path = "/b" // 把请求的URI修改
		r.HandleContext(c)        // 继续后续的处理
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "B",
		})
	})
	r.Run()
}
