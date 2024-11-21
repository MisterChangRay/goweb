package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	db "goweb/models"
	"goweb/pkg/logger"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/calendar"
	beego "github.com/beego/beego/v2/server/web"
)

type MyBizController struct {
	beego.Controller
}

type MsgNotify struct {
	//企业微信的CorpID，当为第三方套件回调事件时，CorpID的内容为suiteid
	ToUserName string
	AgentID    string //接收的应用id，可在应用的设置页面获取
	Encrypt    string
}

type RealMsg struct {
	FromUserName string
	CreateTime   string
	MsgType      string
	Content      string
	MsgId        string
	AgentID      string
}

func (this *MyBizController) WXMessage() {
	var req MsgNotify

	err := this.BindXML(&req)
	if err != nil {
		return
	}
	aeskey, _ := base64.StdEncoding.DecodeString("FVtNfmuOhZahBbEMS2eVyYIEaj8kOTvhHyXvHC5zcvp==")
	iv := aeskey[0:16]

	res := Ase256Decode(req.Encrypt, aeskey, iv)
	lenarr := res[16:20]
	lennum := binary.BigEndian.Uint32(lenarr)
	msgarr := res[20 : 20+lennum]
	var realmsg RealMsg
	err = xml.Unmarshal(msgarr, &realmsg)
	if err != nil {
		return
	}
	logger.Log.Debug("收到消息 %s", realmsg.Content)
	this.dealMsg(realmsg)
	this.Data["json"] = "{}"
	this.ServeJSON()
}

func (this *MyBizController) dealMsg(realmsg RealMsg) {
	if strings.HasPrefix(realmsg.Content, "记账") {
		substring := strings.Split(realmsg.Content, "记账")[1]
		re := regexp.MustCompile(`\d+(\.\d+)?`)
		price := re.FindStringSubmatch(substring)[0]
		remark := strings.Replace(substring, price, "", 1)

		db.MYSQL.Insert(&db.Accounting{
			Uid:         realmsg.FromUserName,
			User:        realmsg.FromUserName,
			Create_time: time.Now(),
			Money:       price,
			Remark:      remark,
		})
	}

	if strings.HasPrefix(realmsg.Content, "添加生日") {
		substring := strings.Split(realmsg.Content, "添加生日")[1]
		re := regexp.MustCompile(`\d+(\.\d+)?`)
		yearmonthday := re.FindStringSubmatch(substring)[0]
		name := strings.Replace(substring, yearmonthday, "", 1)

		// 定义一个时间字符串
		timeStr := yearmonthday + " 15:04:05"
		// 定义一个时间格式
		layout := "20060102 15:04:05"
		// 使用Parse根据格式转换字符串到time.Time类型
		now1, _ := time.Parse(layout, timeStr)

		now := calendar.ByTimestamp(now1.Unix())

		birth, _ := strconv.Atoi(fmt.Sprintf("%02d%02d", now.Lunar.GetMonth(), now.Lunar.GetDay()))
		db.MYSQL.Insert(&db.Members{
			Name:        name,
			Birthday:    birth,
			Create_time: time.Now(),
			Birthyear:   "1992",
		})
	}

}

func (this *MyBizController) WXtestCall() {
	// msg_signature := this.GetString("msg_signature")
	// timestamp := this.GetString("timestamp")
	// nonce := this.GetString("nonce")
	echostr := this.GetString("echostr")
	// echostr, _ := url.QueryUnescape(tmp1)

	aeskey, _ := base64.StdEncoding.DecodeString("FVtNfmuOhZahBbEMS2eVyYIEaj8kOTvhHyXvHC5zcvp==")
	iv := aeskey[0:16]

	res := Ase256Decode(echostr, aeskey, iv)
	lenarr := res[16:20]
	lennum := binary.BigEndian.Uint32(lenarr)
	msgarr := res[20 : 20+lennum]
	msg := string(msgarr)

	logger.Log.Debug("收到消息 %s", msg)

	// var jsonMap map[string]string
	// json.Unmarshal([]byte(res), &jsonMap)

	// var msg string = echostr
	// tmp := []byte(msg)
	this.Ctx.Output.Header("Content-Type", "text/plain;charset=UTF-8")
	this.Ctx.ResponseWriter.Write(msgarr)

}

func Ase256Decode(cipherText string, bKey []byte, bIV []byte) (decryptedString []byte) {
	// bKey := []byte(encKey)
	// bIV := []byte(iv)
	// bKey, _ := base64.StdEncoding.DecodeString(encKey)
	// bIV, _ := base64.StdEncoding.DecodeString(iv)
	// cipherTextDecoded, err := hex.DecodeString(cipherText)
	cipherTextDecoded, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return cipherTextDecoded
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
