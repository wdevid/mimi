package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"html/template"
	"github.com/astaxie/beego/orm"
	"mimi/models"
	"github.com/astaxie/beego/context"
)

type LoginController struct  {
	beego.Controller
}

func (this *LoginController) Get()  {
	tmpl, err := template.ParseFiles("./static/view/login.html")
	if err != nil {
		fmt.Println("Error happened..")
	}
	tmpl.Execute(this.Ctx.ResponseWriter,nil)
	//this.TplName = "./view/login.html"
	//tmpl, err := template.ParseFiles("./view/login.html")
	//if err != nil {
	//	fmt.Println("Error happened..")
	//}
	//type person struct {
	//	Id      int
	//	Name    string
	//	Country string
	//}
	//
	//liumiaocn := person{Id: 1001, Name: "liumiaocn", Country: "static/img/home.jpg"}
	//tmpl.Execute(this.Ctx.ResponseWriter,liumiaocn)
	//
	//name := this.GetString("username","")
	//pd := this.GetString("password","")
	//fmt.Println(name,pd)
}
func (this *LoginController) Post()  {
	o:=orm.NewOrm()
	o.Using("user")
	var user []models.User
	uname := this.Input().Get("username")
	pwd := this.Input().Get("password")
	o.Raw("select user_name,pass_word from user where user_name=? and pass_word=?",uname,pwd).QueryRows(&user)
	fmt.Println(user)
	if len(user)==0 {
		this.Ctx.WriteString("登陆失败")
	}else {
		maxAge := 1<<31-1
		this.Ctx.SetCookie("uname",uname,maxAge,"/")
		this.Ctx.SetCookie("pwd",pwd,maxAge,"/")
		this.Redirect("/",301)
	}

	return ;
}

func CheckLogin(ctx *context.Context) bool {

	o:=orm.NewOrm()
	o.Using("user")
	var user []models.User

	ck,err :=ctx.Request.Cookie("uname")
	if  err != nil{
		return false
	}
	uname := ck.Value
	pk,perr :=ctx.Request.Cookie("pwd")
	if perr != nil {
		return false
	}
	pwd := pk.Value
	o.Raw("select user_name,pass_word from user where user_name=? and pass_word=?",uname,pwd).QueryRows(&user)
	if len(user) == 0 {
		return false
	}else{
		return true
	}
	return false
}

