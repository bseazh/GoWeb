package main

import (
	"Lesson26/dao"
	"Lesson26/models"
	"Lesson26/routers"
)

func main() {
	//	创建数据库
	// create database bubble
	//	连接数据库

	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 程序退出关闭数据库
	defer dao.Close()
	//	模型绑定
	dao.DB.AutoMigrate(&models.Todo{}) //	todos
	// 注册路由
	r := routers.SetupRouter()

	r.Run()

}
