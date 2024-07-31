package service

import (
	"fmt"
	config "goweb/config/db"
	"goweb/service/pojo"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

// Multipart/Urlencoded Form
func GetKey(c *gin.Context) {
	key := c.Query("key")

	res := DoGetKey1(key)
	c.JSON(http.StatusOK, gin.H{
		"value": res,
	})
}

func DoGetKey0(key string) *string {
	u := config.T_keyvalue{Key: key}
	err := config.DB.Read(&u)
	if err == nil {
		return &u.Value
	}
	return nil

}

func DoGetKey1(key string) *string {
	var maps []orm.Params
	num, err := config.DB.Raw("select value from t_keyvalue where `key` = ? and now() < expire_time ", key).Values(&maps)

	if err == nil && num > 0 {
		s := fmt.Sprintln(maps[0]["value"])

		return &s
	}
	return nil
}

func DoDeleteKey0(key string) int64 {
	u := config.T_keyvalue{Key: key}

	num, err := config.DB.Delete(&u)
	if err == nil {
		return num
	}
	return -1
}

// 接受json参数, 及参数校验
func AddKey(c *gin.Context) {
	var req pojo.HelloReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "hello," + req.Name,
	})
}

// 路径参数演示 xxxx.com/:name
func DeleteKey(c *gin.Context) {
	key := c.Query("key")
	num := DoDeleteKey0(key)
	var code string = "0000"
	if num <= 0 {
		code = "9999"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "done",
	})
}
