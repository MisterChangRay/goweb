package common

import (
	pojos "goweb/controllers"
	db "goweb/models"

	beego "github.com/beego/beego/v2/server/web"
)

type CarController struct {
	beego.Controller
}

func (this *CarController) QueryCars() {
	car := this.GetString("car")

	var res pojos.BaseRes
	var lists []db.Mycar
	num, err := db.MYSQL.Raw("select * from mycar where car = ? ", car).QueryRows(&lists)
	if err != nil || num == 0 {
		res = pojos.BaseRes{
			Msg: "no data",
		}
	} else {
		res = pojos.BaseRes{
			Msg:  "okk",
			Data: lists,
		}
	}

	this.Data["json"] = &res
	this.ServeJSON()

}
