package main

import (
	_ "goweb/models"
	_ "goweb/pkg/logger"
	_ "goweb/routers"
	_ "goweb/service"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	beego.Run()
}
