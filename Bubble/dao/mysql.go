package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitMySQL() (err error) {

	DB, err = gorm.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/bubble?charset=utf8&mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		return
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
