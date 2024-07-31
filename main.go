package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"goweb/config"

	"github.com/gin-gonic/gin"
)

func main() {

	// Logging to a file.
	filename := fmt.Sprintf("logs/gin_%s.log", time.Now().Format("2006_01_02_150405"))
	f, _ := os.Create(filename)
	gin.DefaultWriter = io.MultiWriter(f)

	// 1.创建路由
	var r *gin.Engine = gin.Default()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	log.SetOutput(gin.DefaultWriter)
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	config.RegisterRouter(r)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
