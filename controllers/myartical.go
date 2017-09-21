package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"html/template"
	"github.com/astaxie/beego/orm"
	"mimi/models"
)

type MyArticalController struct {
	beego.Controller
}

func (this *MyArticalController) Get() {
	tmpl, err := template.ParseFiles("./static/view/myartical.html")
	if err != nil {
		fmt.Println("Error happened..")
	}
	type imgpath struct {
		Path string
		Cuss [] models.Customer
	}
	o := orm.NewOrm()
	o.Using("user")
	var articals models.User
	uname, _ := this.Ctx.Request.Cookie("uname")
	fmt.Println(o.QueryTable("user").Filter("user_name",uname.Value).One(&articals))

	o.Using("customer")
	var cuss [] models.Customer
	o.QueryTable("customer").OrderBy("-id").Filter("uname",articals.UserName).All(&cuss)
	fmt.Println(cuss)
	if len(cuss) > 0 {
		p := imgpath{Path: "static/img/home.jpg",Cuss:cuss}
		tmpl.Execute(this.Ctx.ResponseWriter, p)
	}else {
		p := imgpath{Path: "static/img/home.jpg"}
		tmpl.Execute(this.Ctx.ResponseWriter, p)
	}

}

