package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//	1.	定义模型
type User struct {
	gorm.Model
	Name   string
	Age    int
	Active bool
}

func main() {
	//	2.	连接数据库
	db, err := gorm.Open("mysql",
		"root:root@(127.0.0.1:3306)/db1?charset=utf8&mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("connect DB failed , err : %v\n", err)
		return
	}
	defer db.Close()
	fmt.Printf("连接数据库成功!\n")

	//	3.	数据库的表与模型建立关系
	db.AutoMigrate(User{})

	//	4.	创建数据添加到库中
	db.Create(&User{Name: "Crazy", Age: 12, Active: false})
	db.Create(&User{Name: "LiMei", Age: 16, Active: true})
	db.Create(&User{Name: "GeLei", Age: 20, Active: false})
	db.Create(&User{Name: "Will", Age: 30, Active: true})

	//	删除
	//	删除记录
	//	警告 删除记录时，请确保主键字段有值，GORM 会通过主键去删除记录，如果主键为空，GORM 会删除该 model 的所有记录。
	//	UPDATE `users` SET `deleted_at`='2021-03-09 10:59:34'  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 3
	var user User
	db.Where("ID = 3 ").First(&user)
	user.Name = ""
	user.Age = 44
	db.Debug().Delete(&user)
	fmt.Printf("%#v\n", user)
	//	user ->	Name:"", Age:44, Active:false
	//	只会根据主键删除其对应的字段，但不会返回对应的值给user

	//	批量删除
	//	删除全部匹配的记录
	//	UPDATE `users` SET `deleted_at`='2021-03-09 11:05:40'  WHERE `users`.`deleted_at` IS NULL AND ((name LIKE '%W%'))
	db.Debug().Where("name LIKE ?", "%W%").Delete(User{})

	//	UPDATE `users` SET `deleted_at`='2021-03-09 11:05:40'  WHERE `users`.`deleted_at` IS NULL AND ((name LIKE '%W%'))
	db.Debug().Delete(User{}, "name LIKE ?", "%W%")

	//	软删除(针对结构体中加载了gorm.Model)
	//	如果一个 model 有 DeletedAt 字段，他将自动获得软删除的功能！
	//	当调用 Delete 方法时， 记录不会真正的从数据库中被删除， 只会将DeletedAt 字段的值会被设置为当前时间

	// Unscoped 方法可以查询被软删除的记录
	//	SELECT * FROM `users`  WHERE (name LIKE '%W%')
	users := User{}
	db.Debug().Unscoped().Where("name LIKE ?", "%W%").Find(&users)
	fmt.Printf("%#v\n", users)

	//	物理删除
	// Unscoped 方法可以物理删除记录
	//	DELETE FROM `users`  WHERE `users`.`id` = 4
	db.Debug().Unscoped().Delete(&users)
}
