package main

import (
	"github.com/astaxie/beego"
	"samples/mimi/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/shorten", &controllers.ShortController{})
	beego.Router("/v1/expand", &controllers.ExpandController{})
	beego.Router("/login", &controllers.LoginControllers{})
	beego.Run()
}
