package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"mimi/models"
)

type DeleteController struct {
	beego.Controller
}

func (this *DeleteController)Get()  {
	var cuss models.Customer
	id,err:=this.GetInt64("id",-1)
	cuss.Id = id
	if err!=nil {
		fmt.Println(err.Error)
	}
	o := orm.NewOrm()
	o.Using("customer")
	o.QueryTable("customer").Filter("id",id).Delete()
	this.Redirect("/myartical",301)
}
