package main

import (
	_ "goweb/models"
	_ "goweb/pkg/logger"
	_ "goweb/routers"
	_ "goweb/service"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// str := "记账 1233 吃饭1266"
	// substring := strings.Split(str, "记账")[1]
	// re := regexp.MustCompile(`\d+(\.\d+)?`)
	// price := re.FindStringSubmatch(str)[0]
	// remark := strings.Replace(substring, price, "", 1)
	// fmt.Printf("%s %s", price, remark)
	beego.Run()
}
