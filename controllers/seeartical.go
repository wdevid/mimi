package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"html/template"
	"mimi/models"
	"github.com/astaxie/beego/orm"
)

type SeeArticalController struct {
	beego.Controller
}

func (this *SeeArticalController)Get()  {
	tmpl, err := template.ParseFiles("./static/view/seeartical.html")
	if err != nil {
		fmt.Println("Error happened..")
	}
	var cuss []models.Customer
	id,err:=this.GetInt64("id",-1)
	if err!=nil {
		 fmt.Println(err.Error)
	}
	o := orm.NewOrm()
	o.Using("customer")
	o.QueryTable("customer").Filter("id",id).All(&cuss)
	type imgpath struct {
		Path string
		Cuss []models.Customer
	}
	fmt.Println("_________=")
	fmt.Println(cuss)
	fmt.Println("_________=")
	p:= imgpath{Path:"static/img/home.jpg",Cuss:cuss}
	tmpl.Execute(this.Ctx.ResponseWriter,p)
}
