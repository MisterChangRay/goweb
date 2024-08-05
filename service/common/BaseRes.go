package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 基本返回包装类
type BaseRes struct {
	Code int
	Msg  string
	Data interface{}
}

// 分页数据
type BasePage struct {
	Data     interface{}
	Total    int
	PageNo   int
	PageSize int
}

func NewBaseRes(data interface{}) BaseRes {
	return BaseRes{
		Code: 0,
		Msg:  "okk",
		Data: data,
	}
}

func HttpSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BaseRes{
		Code: 0,
		Msg:  "okk",
		Data: data,
	})
}

func HttpFail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, BaseRes{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func HttpFail1(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, BaseRes{
		Code: -1,
		Msg:  msg,
		Data: nil,
	})
}

func (m BaseRes) IsSuccess() bool {
	return m.Code == 0
}
