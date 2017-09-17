package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"html/template"
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
	type person struct {
		Id      int
		Name    string
		Country string
		IsLogin bool
	}
	liumiaocn := person{Id: 1001, Name: "liumiaocn", Country: "static/img/home.jpg",IsLogin:CheckLogin(this.Ctx)}
	tmpl.Execute(this.Ctx.ResponseWriter,liumiaocn)
}
