package common

import "time"

type HelloReq struct {
	// 名称必填,长度2-10
	Name       string    `form:"name" json:"name" xml:"name"  binding:"required,max=10,min=2"`
	Account    string    `form:"account" json:"account" xml:"account"  binding:"required"`
	Password   string    `form:"password" json:"password" xml:"password"  binding:"required"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	// 年龄只能 0 -130区间
	Age uint32 `form:"age" json:"age" xml:"age" binding:"gte=0,lte=130"`
}

type KV struct {
	Key   string `form:"key" json:"key" xml:"key"  binding:"required,max=50,min=1"`
	Value string `form:"value" json:"value" xml:"value"  binding:"required,max=1024,min=1"`
	TTL   int32
}
