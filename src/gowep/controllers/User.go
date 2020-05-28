package controllers

import "github.com/astaxie/beego"

type RegController struct {
	beego.Controller
}

func (this *RegController) ShowReg() {
	this.TplName = "register.html"
}

func (this *RegController) HandleReg() {
	name := this.GetString("userName")
	passwd := this.GetString("password")
	if name == "" || passwd == "" {
		beego.Info("账号密码不能为空")
		this.TplName = "regster.html"
		return
	}
}
