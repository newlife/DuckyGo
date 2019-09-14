package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database() {
	connString := "root:Zl@138119@/hello?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open("mysql", connString)
	// 开启Shell显示SQL
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(300)
	//打开
	db.DB().SetMaxOpenConns(500)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	db.SingularTable(true)
	DB = db

}
