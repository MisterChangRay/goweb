package main

import (
	"goweb/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {

	// Simple group: v1
	v1 := engine.Group("/console")
	{
		v1.GET("/hello", service.Hello)

	}

	// Simple group: v2
	// v2 := myRouter.Group("/v2")
	// {
	// 	v2.POST("/login", loginEndpoint)
	// 	v2.POST("/submit", submitEndpoint)
	// 	v2.POST("/read", readEndpoint)
	// }

}
