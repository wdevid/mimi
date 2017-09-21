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
	type imgpath struct {
		Path string
	}
	p:= imgpath{Path:"static/img/home.jpg"}
	tmpl.Execute(this.Ctx.ResponseWriter,p)
}
func (this *LoginController) Post()  {
	o:=orm.NewOrm()
	o.Using("user")
	var user []models.User
	uname := this.Input().Get("username")
	pwd := this.Input().Get("password")
	if uname == ""||pwd == "" {
		this.Ctx.WriteString("用户名或者密码不能为空")
		return
	}
	o.Raw("select user_name,pass_word from user where user_name=? and pass_word=?",uname,pwd).QueryRows(&user)
	fmt.Println(user)
	if len(user)==0 {
		this.Ctx.WriteString("登陆失败")
		this.Redirect("/",301)
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
	cname:=ctx.GetCookie("uname")
	cpwd:=ctx.GetCookie("pwd")
	fmt.Println(cname)
	if len(user) == 0 &&cname!=uname&&cpwd!=pwd{
		return false
	}else{
		return true
	}
	return false
}

