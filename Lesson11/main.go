package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// GET 请求URL?name=`???`&age=`??`
	// "key = value" 多个key=value用"&"连接
	// <=> /index/?name=靓女&age=22

	r.GET("/index", func(c *gin.Context) {
		// 三种方式获取相应字段的值
		//name := c.Query("name")
		//name := c.DefaultQuery("name","靓女")
		name, ok := c.GetQuery("name")
		if !ok {
			name = "靓女"
		}
		// key - value 都为string
		age := c.DefaultQuery("age", "18")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9000")
}
