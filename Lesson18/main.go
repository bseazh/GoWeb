package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Printf("index...\n")
	name, ok := c.Get("name") // 从上下文取值(跨中间件存取值)
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"username": name,
	})

}

// 定义一个中间件m1;统计请求处理函数的耗时
func m1(c *gin.Context) {

	fmt.Printf("m1 in...\n")
	start := time.Now()
	// 计时
	c.Next() // 调用后续的处理函数
	//c.Abort() // 阻止后续的函数执行
	cost := time.Since(start)
	fmt.Printf("cost : %v\n", cost)
	fmt.Printf("m1 out ... \n")
}

// 定义m2的中间件来测试
func m2(c *gin.Context) {
	c.Set("name", "靓仔")
	fmt.Printf("m2 in ...\n")
	c.Next()
	fmt.Printf("m2 out ...\n")
}

func authMiddleware(doCheck bool) func(*gin.Context) {
	// 连接数据库
	// 或者做一些其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			// 存放具体的逻辑
			// 是否登录的判断
			// if 是登录用户
			c.Next()
			// else 不是登录用户
			//		c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default() // Logger 和 Recovery 中间件
	// 其中Logger用于写日志
	// Recovery用于错误处理
	r.Use(m1, m2, authMiddleware(true))
	// GET(relativePath string , handlers... HandlerFunc ) IRoutes
	r.GET("/index", indexHandler)
	r.GET("/user", indexHandler)
	r.GET("/shop", indexHandler)

	//路由组注册中间件  方法1:
	//xxGroup := r.Group("/xx",authMiddleware(true))
	//{
	//	xxGroup.GET("/index", func(c *gin.Context) {
	//		c.JSON(http.StatusOK,gin.H{
	//			"msg":"/xx/index",
	//		})
	//	})
	//}
	////路由组注册中间件  方法2:
	//xx2Group := r.Group("/xx2")
	//xx2Group.Use(authMiddleware(true))
	//{
	//	xx2Group.GET("/index", func(c *gin.Context) {
	//		c.JSON(http.StatusOK,gin.H{
	//			"msg":"/xx2/index",
	//		})
	//	})
	//}
	r.Run()
}
