package main

import (
	"DuckyGo/model"
	"DuckyGo/server"
)

func main() {
	// 装载路由
	model.Database()
	r := server.NewRouter()

	// 运行 起在8000端口
	r.Run(":8000")
}
