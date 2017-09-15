package main

import (
	"github.com/astaxie/beego"
	"samples/mimi/controllers"
	"github.com/astaxie/beego/orm"
	"samples/mimi/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	models.RegisterDB()
	//// 开启 ORM 调试模式
	orm.Debug = true
	//// 自动建表
	orm.RunSyncdb("default", false, true)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/regester", &controllers.RegesterController{})
	beego.Run()
}
