package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义结构体
// 命名规则 以大写为间隔分开,大写变小写,最后加's'
type UseInfo struct {
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 命令行打开mysql, "mysql -u root -p"
	db, err := gorm.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/db1?charset=utf8&mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("")
		panic(err)
	}
	defer db.Close()

	fmt.Printf("成功连接数据库!\n")
	//创建表: 自动迁移
	db.AutoMigrate(UseInfo{})
	// 添加行
	user := UseInfo{
		Name:   "靓仔",
		Gender: "男",
		Hobby:  "媾女",
	}
	db.Create(&user)

	// 查询
	db.First(&user)
	fmt.Printf("%#v\n", user)

	// 更新
	db.Model(&user).Update("hobby", "泡妞")
	fmt.Printf("%#v\n", user)

	// 删除
	db.Delete(&user)
}
