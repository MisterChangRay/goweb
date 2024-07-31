package config

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var DB orm.Ormer = nil

// Model Struct
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

func InitOrm() {
	// register model
	orm.RegisterModel(new(T_user))
	orm.RegisterModel(new(T_keyvalue))
	orm.Debug = true

	// set default database
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dburi"), 30)

	// create table
	orm.RunSyncdb("default", false, true)
	DB = orm.NewOrm()

}
