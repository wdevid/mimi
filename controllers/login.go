package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"html/template"
)

type LoginControllers struct  {
	beego.Controller
}

func (this *LoginControllers) GET()  {

	tmpl, err := template.ParseFiles("./view/login.html")
	if err != nil {
		fmt.Println("Error happened..")
	}
	type person struct {
		Id      int
		Name    string
		Country string
	}

	liumiaocn := person{Id: 1001, Name: "liumiaocn", Country: "static/img/home.jpg"}
	tmpl.Execute(this.Ctx.ResponseWriter,liumiaocn)

	name := this.GetString("username","")
	pd := this.GetString("password","")
	fmt.Println(name,pd)
}

