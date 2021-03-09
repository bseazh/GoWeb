package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pwd"`
}

func main() {
	r := gin.Default()

	// './'是当前目录别忘了加'.'
	r.LoadHTMLFiles("./index.html")
	// 在"/user"路径下,GET请求
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := UserInfo{
		//	Username : username ,
		//	Password : password ,
		//}
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"Message": "OK",
			})
		}
	})

	// 在"/index"路径下,GET请求
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 在html中声明：<form action="/form" method="post">

	// 就会在form页面下提出请求"post"
	// 页面返回OK，在控制台中打印绑定的参数
	r.POST("/form", func(c *gin.Context) {
		var user UserInfo
		if err := c.ShouldBind(&user); err == nil {
			fmt.Printf("username : %s , password : %s\n", user.Username, user.Password)
			c.JSON(http.StatusOK, gin.H{
				"message": "Ok",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	// 利用postman调试
	r.GET("/json", func(c *gin.Context) {
		var user UserInfo
		if err := c.ShouldBind(&user); err == nil {
			fmt.Printf("username : %s , password : %s\n", user.Username, user.Password)
			c.JSON(http.StatusOK, gin.H{
				"message": "Ok",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})
	r.Run(":9090")
}
