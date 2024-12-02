package controllers

import (
	"bytes"
	"fmt"
	_ "goweb/models"
	db "goweb/models"
	_ "goweb/pkg/logger"
	_ "goweb/routers"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/calendar"
	"github.com/jasonlvhit/gocron"
)

func init() {
	s := gocron.NewScheduler()
	s.Every(1).Days().At("09:00:00").Do(birthday)
	s.Start()
}

// 扫描生日并推送
func birthday() {
	t := time.Now()
	// 1. ByTimestamp
	// 时间戳
	now := calendar.ByTimestamp(t.Unix())
	md1, _ := strconv.Atoi(fmt.Sprintf("%02d", int64(now.Lunar.GetMonth())) + fmt.Sprintf("%02d", int64(now.Lunar.GetDay())))
	now2 := calendar.ByTimestamp(t.Unix() + int64(7*24*60*60))
	md2, _ := strconv.Atoi(fmt.Sprintf("%02d", int64(now2.Lunar.GetMonth())) + fmt.Sprintf("%02d", int64(now2.Lunar.GetDay())))

	var lists []db.Members

	num, err := db.MYSQL.Raw("select * from members where `birthday` >= ? and `birthday` <= ? ", md1, md2).QueryRows(&lists)
	if err != nil || num == 0 {
		return
	}

	msg := ""
	nowmsg := ""
	for _, v := range lists {
		if v.Birthday == md1 {
			nowmsg += fmt.Sprintf("%s 今天生日!\r\n", v.Name)
		} else {
			msg += fmt.Sprintf("%s 还有 %d 天生日!\r\n", v.Name, v.Birthday-md1)

		}
	}
	msg = nowmsg + msg

	// logger.Log.Debug("收到请求 %v", msg)
	sendQun(msg)
}

func sendQun(contet string) {
	// 发送群消息
	tmp := fmt.Sprintf(`{"msgtype": "text","text": {"content": "%s"}}`, contet)

	http.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=ac533940-9698-4622-a59b-bc9f880270a4", "application/json", bytes.NewBuffer([]byte(tmp)))

}

func sendMy(content string) {
	requestURL := fmt.Sprintf("http://47.109.108.16:7517/wecomchan?sendkey=Wkiren23714_JJs&msg=%s&msg_type=text", url.QueryEscape(content))
	http.Get(requestURL)

}
