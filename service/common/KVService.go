package common

import (
	"fmt"
	db "goweb/config/db"
	"net/http"
	"time"

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
	u := db.T_keyvalue{Key: key}
	err := db.DB.Read(&u)
	if err == nil {
		return &u.Value
	}
	return nil

}

func DoGetKey1(key string) *string {
	var maps []orm.Params
	num, err := db.DB.Raw("select value from t_keyvalue where `key` = ? and now() < expire_time ", key).Values(&maps)

	if err == nil && num > 0 {
		s := fmt.Sprintln(maps[0]["value"])

		return &s
	}
	return nil
}

func DoDeleteKey0(key string) int64 {
	u := db.T_keyvalue{Key: key}

	num, err := db.DB.Delete(&u)
	if err == nil {
		return num
	}
	return -1
}

// Multipart/Urlencoded Form
func UpdateKey(c *gin.Context) {
	var req KV
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var expireTime time.Time
	if req.TTL > 0 {
		expireTime = time.Now().Add(time.Second * time.Duration(req.TTL))
	}

	res := DoUpdateKey0(req.Key, req.Value, expireTime)
	c.JSON(http.StatusOK, gin.H{
		"value": res,
	})
}

func DoUpdateKey0(key string, value string, time time.Time) int64 {
	if time.IsZero() {
		_, err := db.DB.Raw("update t_keyvalue set value = ? , `update_time` = now()  where `key` = ? and now() < expire_time ", value, key).Exec()
		if err == nil {
			return 0
		}
	} else {
		_, err := db.DB.Raw("update t_keyvalue set value = ? , `update_time` = now() ,expire_time=? where `key` = ? and now() < expire_time ", value, time, key).Exec()
		if err == nil {
			return 0
		}
	}

	return -1
}

// 接受json参数, 及参数校验
func AddKey(c *gin.Context) {
	var req KV
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := DoAddOrUpdateKey(req)
	var code string = "9999"
	if res == 0 {
		code = "0000"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "okk",
	})
}

func DoAddOrUpdateKey(req KV) int32 {
	ins := db.T_keyvalue{
		Key:         req.Key,
		Value:       req.Value,
		Create_time: time.Now(),
		Update_time: time.Now(),
		Expire_time: time.Now().Add(time.Second * time.Duration(req.TTL)),
	}
	_, err := db.DB.Insert(&ins)
	if err == nil {
		return 0
	}
	return -1
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
