package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"html/template"
	"github.com/astaxie/beego/orm"
	"mimi/models"
)

type SearchController struct {
	beego.Controller
}

func (this *SearchController)Get()  {
	fmt.Println(CheckLogin(this.Ctx))
	tmpl, err := template.ParseFiles("./static/view/home.html")
	if err != nil {
		fmt.Println("Error happened..")
	}
	con := this.Input().Get("search")
	o := orm.NewOrm()
	o.Using("customer")
	var cuss []models.Customer
	o.QueryTable("customer").Filter("content__contains",con).All(&cuss)
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
