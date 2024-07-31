package config

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var DB orm.Ormer = nil

// Model Struct
type User struct {
	Id   int    `orm:"pk;auto"`
	Name string `orm:"size(100)"`
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
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(T_keyvalue))
	orm.Debug = true

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:2vKeG&1.3@tcp(47.109.108.16:7501)/fastapi_demo?charset=utf8&timeout=6s&readTimeout=15s&interpolateParams=true&parseTime=true&loc=Local", 30)

	// create table
	orm.RunSyncdb("default", false, true)
	DB = orm.NewOrm()

}
