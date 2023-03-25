package main

import (
	"blog/model"
	"blog/routes"
)

// 先建立数据结构

func main() {
	model.InitDb()

	routes.InitRouter()
}
