package controllers

import "github.com/astaxie/beego"

type LoginOutController struct {
	beego.Controller
}

func (this *LoginOutController)Get()  {
	uname := this.GetSession("uname")
	if uname!="未登录" {
		maiAge := 1<<31 - 1
		this.Ctx.SetCookie("uname","未登录",maiAge,"/")
		this.Redirect("/",301)
	}
}
