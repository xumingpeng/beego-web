package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	UserName string
	passwd   string
}

func init() {
	//连接数据库
	err := orm.RegisterDataBase("default", "mysql", "homestead:server@tcp(192.168.10.10:3306)/hxxd")
	if err != nil {
		beego.Info("连接数据库失败", err)
		return
	}
	//注册模型
	orm.RegisterModel(new(User))
	//自动创建表 参数二为是否开启创建表 参数三为是否更新表
	err2 := orm.RunSyncdb("default", false, true)
	if err2 != nil {
		beego.Info("orm.RunSyncdb err", err)
		return
	}
}
