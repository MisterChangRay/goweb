package common

import (
	"encoding/json"
	pojos "goweb/controllers"
	"goweb/pkg/logger"

	beego "github.com/beego/beego/v2/server/web"
)

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Formdata() {
	var req HelloReq

	err := this.BindForm(&req)
	if err == nil {
		jsondata, _ := json.Marshal(req)
		logger.Log.Debug("收到请求 %s", string(jsondata))
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
