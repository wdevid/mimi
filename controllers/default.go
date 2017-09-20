package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"html/template"
	"github.com/astaxie/beego/orm"
	"mimi/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	fmt.Println(CheckLogin(this.Ctx))
	tmpl, err := template.ParseFiles("./static/view/home.html")
	if err != nil {
		fmt.Println("Error happened..")
	}
	o := orm.NewOrm()
	o.Using("customer")
	var cuss []models.Customer
	o.QueryTable("customer").OrderBy("-id").All(&cuss)
	type person struct {
		Id      int
		Name    string
		Country string
		IsLogin bool
		Cuss []models.Customer
	}
	fmt.Println(cuss,this.Ctx.GetCookie("uname"))
	liumiaocn := person{Id: 1001, Name: this.Ctx.GetCookie("uname"), Country: "static/img/home.jpg",IsLogin:CheckLogin(this.Ctx),Cuss:cuss}


	tmpl.Execute(this.Ctx.ResponseWriter,liumiaocn)
}


