package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var MYSQL orm.Ormer = nil

// 数据库表单定义
type T_user struct {
	Id          int    `orm:"pk;auto"`
	Name        string `orm:"size(100)"`
	status      int16
	Create_time time.Time
	Update_time time.Time
}

type T_keyvalue struct {
	Key         string `orm:"pk;size(255)"`
	Value       string `orm:"size(1024)"`
	Create_time time.Time
	Update_time time.Time
	Expire_time time.Time
}

type Members struct {
	Name        string `orm:"pk;size(255)"`
	Phone       string `orm:"size(255)"`
	Birthday    int    `orm:"int"`
	Birthyear   string `orm:"int"`
	Create_time time.Time
}

type Accounting struct {
	Id          uint32 `pk;orm:"int"`
	Uid         string `orm:"size(255)"`
	Remark      string `orm:"size(255)"`
	Money       string `orm:"decimal"`
	Create_time time.Time
	User        string `orm:"size(255)"`
}

// 初始化数据库相关代码
func init() {
	// register model
	orm.RegisterModel(new(T_user))
	orm.RegisterModel(new(T_keyvalue))
	orm.RegisterModel(new(Members))
	orm.RegisterModel(new(Accounting))
	orm.Debug = true

	dburl, err := beego.AppConfig.String("dburi")
	if err == nil {
		// set default database
		orm.RegisterDataBase("default", "mysql", dburl)

		// create table
		orm.RunSyncdb("default", false, true)
		MYSQL = orm.NewOrm()

	}

}
