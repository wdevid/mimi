package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"html/template"
	"github.com/astaxie/beego/orm"
	"mimi/models"
)

type RegesterController struct {
	beego.Controller
}

func (this *RegesterController) Get()  {
	tmpl, err := template.ParseFiles("./static/view/regester.html")
	if err != nil {
		fmt.Println("Error happened..")
	}
	type imgpath struct {
		Path string
	}
	p:= imgpath{Path:"static/img/home.jpg"}
	tmpl.Execute(this.Ctx.ResponseWriter,p)

}

func (this *RegesterController) Post() {
	u := new(models.User)
	o := orm.NewOrm()
	o.Using("user")
	uname := this.Input().Get("username")
	pwd := this.Input().Get("password")
	u.UserName = uname
	u.PassWord = pwd
	fmt.Println(o.Insert(u))
	this.Redirect("/login",301)
	return ;
}

