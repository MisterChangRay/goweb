package common

import (
	pojos "goweb2/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Formdata() {
	var req HelloReq

	err := this.BindForm(&req)
	if err == nil {
		res := pojos.BaseRes{
			Msg: req.Name,
		}

		this.Data["json"] = &res
		this.ServeJSON()
	}

}

func (this *HelloController) Jsondata() {
	var req HelloReq

	err := this.BindJSON(&req)
	if err == nil {
		res := pojos.BaseRes{
			Msg: req.Name,
		}

		this.Data["json"] = &res
		this.ServeJSON()
	}
}

func (this *HelloController) PathParam() {
	name := this.GetString(":name")
	res := pojos.BaseRes{
		Msg: name,
	}

	this.Data["json"] = &res
	this.ServeJSON()
}

func (this *HelloController) Queryparam() {
	name := this.GetString("name")
	res := pojos.BaseRes{
		Msg: name,
	}

	this.Data["json"] = &res
	this.ServeJSON()

}
