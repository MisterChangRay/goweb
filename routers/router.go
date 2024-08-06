package routers

import (
	"goweb/controllers"
	"goweb/controllers/common"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	v1api := beego.NewNamespace("/v1/api",
		beego.NSNamespace("/hello",
			beego.NSRouter("/formdata", &common.HelloController{}, "post:Formdata"),
			beego.NSRouter("/jsondata", &common.HelloController{}, "post:Jsondata"),
			beego.NSRouter("/pathparam/:name", &common.HelloController{}, "get,post:PathParam"),
			beego.NSRouter("/queryparam", &common.HelloController{}, "get,post:Queryparam"),
		),
		beego.NSRouter("/kvsetting", &common.KVController{}),
	)
	beego.AddNamespace(v1api)
}
