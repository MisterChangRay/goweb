package common

import (
	"context"
	"fmt"
	pojos "goweb/controllers"
	db "goweb/models"
	"time"

	"github.com/beego/beego/v2/client/orm"

	beego "github.com/beego/beego/v2/server/web"
)

type KVController struct {
	beego.Controller
}

func (this *KVController) Get() {
	key := this.GetString("key")
	if key != "" {
		keyvalue := DoGetKey1(key)

		res := pojos.BaseRes{
			Msg: fmt.Sprintf("key %s = %s", key, *keyvalue),
		}

		this.Data["json"] = &res
		this.ServeJSON()
	}

}
func (this *KVController) Post() {
	var req KV

	err := this.BindJSON(&req)
	if err == nil {
		err := db.MYSQL.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
			var expireTime time.Time
			if req.TTL > 0 {
				expireTime = time.Now().Add(time.Second * time.Duration(req.TTL))
			}
			i, err := DoUpdateKey0(req.Key, req.Value, expireTime)
			if i == 0 {
				_, err = DoAddKey(req)
			}
			return err
		})

		var res pojos.BaseRes
		if err == nil {
			res = pojos.BaseRes{
				Msg: fmt.Sprintf("key %s = %s", req.Key, req.Value),
			}
		} else {
			res = pojos.BaseRes{
				Msg: fmt.Sprintf("update key %s failed", req.Key),
			}
		}

		this.Data["json"] = &res
		this.ServeJSON()
	}
}

func (this *KVController) Delete() {
	var req KV

	err := this.BindJSON(&req)
	if err == nil {
		DoDeleteKey0(req.Key)

		res := pojos.BaseRes{
			Msg: fmt.Sprintf("delete key %s success", req.Key),
		}

		this.Data["json"] = &res
		this.ServeJSON()
	}

}

func DoGetKey0(key string) *string {
	u := db.T_keyvalue{Key: key}
	err := db.MYSQL.Read(&u)
	if err == nil {
		return &u.Value
	}
	return nil

}

func DoGetKey1(key string) *string {
	var maps []orm.Params
	num, err := db.MYSQL.Raw("select value from t_keyvalue where `key` = ? and now() < expire_time ", key).Values(&maps)

	var s string = ""
	if err == nil && num > 0 {
		s = fmt.Sprintln(maps[0]["value"])

		return &s
	}
	return &s
}

func DoDeleteKey0(key string) int64 {
	u := db.T_keyvalue{Key: key}

	num, err := db.MYSQL.Delete(&u)
	if err == nil {
		return num
	}
	return -1
}

func DoUpdateKey0(key string, value string, time time.Time) (int64, error) {
	if time.IsZero() {
		_, err := db.MYSQL.Raw("update t_keyvalue set value = ? , `update_time` = now()  where `key` = ? and now() < expire_time ", value, key).Exec()
		if err == nil {
			return 1, err
		}
	} else {
		_, err := db.MYSQL.Raw("update t_keyvalue set value = ? , `update_time` = now() ,expire_time=? where `key` = ? and now() < expire_time ", value, time, key).Exec()
		if err == nil {
			return 1, err
		}
	}

	return -1, nil
}

func DoAddKey(req KV) (int32, error) {
	ins := db.T_keyvalue{
		Key:         req.Key,
		Value:       req.Value,
		Create_time: time.Now(),
		Update_time: time.Now(),
		Expire_time: time.Now().Add(time.Second * time.Duration(req.TTL)),
	}
	_, err := db.MYSQL.Insert(&ins)
	if err == nil {
		return 0, err
	}
	return -1, err
}
