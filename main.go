package main

import (
	_ "goweb/models"
	_ "goweb/pkg/logger"
	_ "goweb/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
