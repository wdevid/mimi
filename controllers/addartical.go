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
		ck,err :=this.Ctx.Request.Cookie("uname")
		if  err != nil{
		}
		uname := ck.Value

		o.Raw("select Id from user where user_name=?",uname).QueryRows(&user)
		cus := new(models.Customer)
		cus.Content = this.Input().Get("content")
		cus.Title = this.Input().Get("title")
		fmt.Println(user)
		cus.Uid = user[0].Id
		cus.Uname = uname
		cus.Created = time.Now()
		cus.Updated = time.Now()
		cus.ReplyTime = time.Now()
		o.Using("customer")
		o.Insert(cus)
		this.Redirect("/",301)
	}else {
		this.Redirect("/login",301)
	}
}


