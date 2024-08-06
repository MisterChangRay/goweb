package main

import (
	"goweb/models"
	_ "goweb/models"
	_ "goweb/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	models.InitOrm()
	beego.Run()
}
