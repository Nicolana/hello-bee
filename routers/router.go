package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"hello/controllers"
	"net/http"
)

var success = []byte("SUPPORT OPTIONS")

var corsFunc = func(ctx *context.Context) {
	origin := ctx.Input.Header("Origin")
	ctx.Output.Header("Access-Control-Allow-Methods", "OPTIONS,DELETE,POST,GET,PUT,PATCH")
	ctx.Output.Header("Access-Control-Max-Age", "3600")
	ctx.Output.Header("Access-Control-Allow-Headers", "X-Custom-Header,accept,Content-Type,Access-Token,Token")
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	ctx.Output.Header("Access-Control-Allow-Origin", origin)
	if ctx.Input.Method() == http.MethodOptions {
		// options请求，返回200
		ctx.Output.SetStatus(http.StatusOK)
		_ = ctx.Output.Body(success)
	}
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, corsFunc)

    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello/:username([\\w]+)", &controllers.HelloController{})
    beego.Get("/alice", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Bob!"))
	})
    beego.Post("/alice", func(ctx *context.Context) {
		ctx.Output.Body([]byte("POST Bob!"))
	})
    ns := beego.NewNamespace("/api",
    	beego.NSNamespace("/user",
			beego.NSRouter("/", &controllers.UserController{}, "get:GetOne"),
			beego.NSRouter("/:id", &controllers.UserController{}, "get:GetOne"),
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login;get:GetLogin"),
		),
	)
    beego.AddNamespace(ns)
}

