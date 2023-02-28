package GoDemofather

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

var Host = "127.0.0.1"
var Port = 3306
var Username = "root"
var Password = "Cm@336699"
var Database = "fzys_v2"

type User struct {
	ID        uint   `gorm:"primary_key"`
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Email     string `gorm:"type:varchar(100);unique_index"`
}

func Test1(t *testing.T) {
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Host, Port, Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDb, err := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	// 自动迁移表结构
	db.AutoMigrate(&User{})

	// 创建用户
	user := User{FirstName: "John", LastName: "Doe", Email: "john@example.com"}
	db.Create(&user)

	// 查询用户
	var result User
	db.First(&result, user.ID)
	fmt.Println(result.FirstName)
}

func Test2(t *testing.T) {
}
