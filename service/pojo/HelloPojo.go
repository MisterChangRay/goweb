package pojo

import "time"

type HelloReq struct {
	Name       string    `form:"user" json:"user" xml:"user"  binding:"required" max:"10" min:"1"`
	Account    string    `form:"account" json:"account" xml:"account"  binding:"required"`
	Password   string    `form:"password" json:"password" xml:"password"  binding:"required"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
}
