package main

import (
	"blog/model"
	"blog/routes"
)

// 先建立数据结构

func main() {
	// 初始化数据库
	model.InitDb()

	// 初始化路由设置
	routes.InitRouter()
}
