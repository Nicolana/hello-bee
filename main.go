package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "hello/routers"
	"hello/utils"
)

func main() {
	utils.InitSql()
	beego.Run()
}

