package routers

import (
	"Lesson26/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() (r *gin.Engine) {
	r = gin.Default()
	//	告诉驱动在哪找static文件
	r.Static("./static", "static")
	//	导入 html文件进行
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", controller.IndexHandle)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodolist)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return
}
