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

	// // 定义一个时间字符串
	// timeStr := "2024-01-06 15:04:05"
	// // 定义一个时间格式
	// layout := "2006-01-02 15:04:05"
	// // 使用Parse根据格式转换字符串到time.Time类型
	// now1, _ := time.Parse(layout, timeStr)

	// now := calendar.ByTimestamp(now1.Unix())
	// fmt.Printf("%s  ", fmt.Sprintf("%02d%02d", now.Lunar.GetMonth(), now.Lunar.GetDay()))

	beego.Run()
}
