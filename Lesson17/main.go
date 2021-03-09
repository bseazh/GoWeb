package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 请求方法的大集合 / 大杂烩
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"Method": "Get",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"Method": "Post",
			})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{
				"Method": "Delete",
			})
		}
	})
	// NoRoute
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Notfound",
		})
	})
	// 路由组的组 多用于区分不同业务线或API
	// 把公用的前缀提取出来,创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Message": "/video/index",
			})
		})
		videoGroup.GET("home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Message": "/video/home",
			})
		})
		videoGroup.GET("user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Message": "/video/user",
			})
		})
	}
	r.Run()
}
