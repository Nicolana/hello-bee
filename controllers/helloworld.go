package controllers

import beego "github.com/beego/beego/v2/server/web"

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Get() {
	var username = this.Ctx.Input.Param(":username")
	this.Ctx.WriteString("Hello " + username)
}
