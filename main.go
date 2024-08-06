package main

import (
	"goweb2/models"
	_ "goweb2/models"
	_ "goweb2/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	models.InitOrm()
	beego.Run()
}
