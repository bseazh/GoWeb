package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64  `gorm:"primary_key"`
	Name string `gorm:"default:'靓仔'"`
	Age  int64  `gorm:"default:22"`
}

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql",
		"root:root@(127.0.0.1:3306)/db1?charset=utf8&mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()
	fmt.Printf("连接数据库成功！\n")

	// 创建表(table)
	db.AutoMigrate(&User{})
	fmt.Printf("创建表成功！\n")

	// 添加元素
	// 01 - 05 : Charles,Mark,Bill,Vincent,William,
	// 06 - 10 : Jseph,James,Henry,Gary,Martin
	users := []User{
		{Name: "Charles", Age: 18},
		{Name: "Mark", Age: 40},
		{Name: "Bill", Age: 54},
		{Name: "Vincent", Age: 17},
		{Name: "William", Age: 59},
		{Name: "Jseph", Age: 50},
		{Name: "James", Age: 30},
		{Name: "Gary", Age: 17},
		{Name: "Henry", Age: 17},
		{Name: "Martin", Age: 95},
	}
	for _, user := range users {
		db.Debug().Create(&user)
	}
	//	高级查询
	//	SELECT id, name, age FROM `users`  WHERE (name = 'Henry')
	var result User
	db.Debug().Table("users").Select("id, name, age").Where("name = ?", "Henry").Scan(&result)
	fmt.Printf("%#v\n", result)
	//	User{ID:9, Name:"Henry", Age:17}

	//	SELECT id, name, age FROM `users`  WHERE (id > 7)
	var results []User
	db.Debug().Table("users").Select("id, name, age").Where("id > ?", 7).Scan(&results)
	fmt.Printf("%#v\n", results)
	//	User{ID:8, Name:"Gary", Age:17}
	//	User{ID:9, Name:"Henry", Age:17}
	//	User{ID:10, Name:"Martin", Age:95}

	//	原生 SQL
	//	SELECT id, name, age FROM users WHERE name = 'Vincent'
	result = User{}
	db.Debug().Raw("SELECT id, name, age FROM users WHERE name = ?", "Vincent").Scan(&result)
	fmt.Printf("%#v\n", result)
	//	User{ID:4, Name:"Vincent", Age:17}
}
