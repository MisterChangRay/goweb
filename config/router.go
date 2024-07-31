package config

import (
	service "goweb/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {

	// Simple group: v1
	v1 := engine.Group("/hello")
	{
		v1.GET("/", service.Hello)
		v1.POST("/formdata", service.Hello_formdata)
		v1.POST("/jsondata", service.Hello_jsondata)
		v1.GET("/pathparam/:name", service.Hello_pathparam)
		v1.GET("/queryparam", service.Hello_queryparam)

	}

	// Simple group: v2
	// v2 := myRouter.Group("/v2")
	// {
	// 	v2.POST("/login", loginEndpoint)
	// 	v2.POST("/submit", submitEndpoint)
	// 	v2.POST("/read", readEndpoint)
	// }

}
