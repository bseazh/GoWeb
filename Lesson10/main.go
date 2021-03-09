package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/Json", func(c *gin.Context) {
		// 方法一：
		//data := map[string]interface{}{
		//	"name" : "靓仔" ,
		//	"msg" : "hello Go" ,
		//	"age" : 18 ,
		//}
		// gin.H 底层逻辑就是map[string]interface{}
		data := gin.H{
			"name": "靓仔",
			"msg":  "hello world",
			"age":  18,
		}
		c.JSON(http.StatusOK, data)
	})
	// 方法2 : 结构体 (字段必须大写)
	// 灵活使用tag来对结构体字段做定制化操作
	type msg struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}

	r.GET("/anotherJson", func(c *gin.Context) {
		data := msg{
			Name:    "靓仔",
			Message: "hello world",
			Age:     18,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run()
}
