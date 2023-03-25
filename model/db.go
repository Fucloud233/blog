package model

import (
	"blog/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	// grom 文档: https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	// mysql 配置: https://github.com/go-sql-driver/mysql#dsn-data-source-name
	// data source name

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("数据库连接失败!\n", err)
	}

	// 会自动迁移schema
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	// 获取通用数据对象然后使用其提供的功能
	sqlDB, _ := db.DB()

	// 设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
