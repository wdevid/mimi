package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mimi/models"
	"time"
	"fmt"
	"html/template"
)

type AddArticalController struct {
	beego.Controller
}

func (this *AddArticalController) Get()  {
	tmpl, err := template.ParseFiles("./static/view/artical.html")
	if err != nil {
		fmt.Println("Error happened..")
	}
	type imgpath struct {
		Path string
	}
	p:= imgpath{Path:"static/img/home.jpg"}
	tmpl.Execute(this.Ctx.ResponseWriter,p)

}
func (this *AddArticalController) Post() {
	if CheckLogin(this.Ctx) {
		o:=orm.NewOrm()
		o.Using("user")
		var user []models.User
		uname := this.GetSession("uname")
		o.Raw("select Id from user where user_name=?",uname).QueryRows(&user)
		cus := new(models.Customer)
		cus.Content = this.Input().Get("content")
		cus.Title = this.Input().Get("title")
		cus.Uid = user[0].Id
		cus.Created = time.Now()
		o.Using("customer")
		o.Insert(cus)
	}else {
		this.Redirect("/login",301)
	}
}
