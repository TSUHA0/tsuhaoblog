package main

import (
	"tsuhaoblog/model"
	"tsuhaoblog/routes"
)

func main() {
	//引用数据库
	model.InitDB()
	routes.InitRouter()
}
