package routers

import (
	"goweb/controllers"
	"goweb/controllers/common"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	beego.Router("/", &controllers.MainController{})

	v1api := beego.NewNamespace("/v1/api",
		beego.NSNamespace("/hello",
			beego.NSRouter("/formdata", &common.HelloController{}, "post:Formdata"),
			beego.NSRouter("/jsondata", &common.HelloController{}, "post:Jsondata"),
			beego.NSRouter("/pathparam/:name", &common.HelloController{}, "get,post:PathParam"),
			beego.NSRouter("/queryparam", &common.HelloController{}, "get,post:Queryparam"),
		),
		beego.NSNamespace("/mypro",
			beego.NSRouter("/wxtestCall", &common.MyBizController{}, "get:WXtestCall"),
			beego.NSRouter("/wxtestCall", &common.MyBizController{}, "post:WXMessage"),
		),
		beego.NSNamespace("/mycar",
			beego.NSRouter("/list", &common.CarController{}, "get:QueryCars"),
		),
		beego.NSRouter("/kvsetting", &common.KVController{}),
	)

	beego.AddNamespace(v1api)

}
