package main

// Mac 命令行打开MySQL: mysql -u root -p
import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	/*
		ID        uint `gorm:"primary_key"` // ID 是gorm的主键
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time `sql:"index"`
	*/
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

type User2 struct {
	ID        uint      // column name is `id`
	Name      string    // column name is `name`
	Birthday  time.Time // column name is `birthday`
	CreatedAt time.Time // column name is `created_at`
}

// 使用`AnimalID`作为主键
type Animal struct {
	AnimalId int64     `gorm:"column:beast_id"`         // set column name to `beast_id`
	Birthday time.Time `gorm:"column:day_of_the_beast"` // set column name to `day_of_the_beast`
	Age      int64     `gorm:"column:age_of_the_beast"` // set column name to `age_of_the_beast`
}

// TableName 是自定义的不受默认规则限制 ; 不受限制，复数，前缀等规则限制
func (User2) TableName() string {
	return "profileBs"
}

func (u User) TableName() string {
	if u.Role == "admin" {
		return "admin_users"
	} else {
		return "userAs"
	}
}
func main() {
	db, err := gorm.Open("mysql",
		"root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//GORM还支持更改默认表名称规则：
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "SB_" + defaultTableName
	}

	//禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)

	// 自动迁移

	db.AutoMigrate(&User{})
	db.AutoMigrate(&User2{})
	db.AutoMigrate(&Animal{})

	var user = User{
		Name: "靓仔",
		Age: sql.NullInt64{
			int64(12),
			false,
		},
		MemberNumber: new(string),
	}

	db.Create(&user) // `CreatedAt`将会是当前时间

	// 可以使用`Update`方法来改变`CreateAt`的值
	db.Model(&user).Update("CreatedAt", time.Now())

	db.Save(&user)                           // `UpdatedAt`将会是当前时间
	db.Model(&user).Update("Name", "jinzhu") // `UpdatedAt`将会是当前时间
}
