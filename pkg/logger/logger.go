package logger

import (
	"github.com/beego/beego/v2/core/logs"
)

var Log *logs.BeeLogger

func init() {

	Log = logs.NewLogger()
	Log.EnableFuncCallDepth(true)
	Log.SetLogFuncCallDepth(2)
	Log.SetLogger(logs.AdapterFile, `{"filename":"./logs/my.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":30,"color":true}`)
	Log.SetLogger(logs.AdapterConsole)
}
