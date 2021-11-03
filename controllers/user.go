package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"hello/models"
	"hello/utils"
	"strconv"
	"time"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) GetOne() {
	token := this.Ctx.Input.Header("Token")
	idStr := this.Ctx.Input.Param(":id")
	et := utils.EasyToken{}
	valid, err := et.ValidateToken(token)
	if !valid {
		this.Ctx.ResponseWriter.WriteHeader(200)
		this.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		this.ServeJSON()
		return
	}
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUserById(id)
	if v == nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = v
	}
	this.ServeJSON()
}

func (this *UserController) Login() {
	var reqData struct {
		Username string `valid:"Required"`
		Password string `valid:"Required"`
	}
	var token string
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &reqData); err == nil {
		if ok, user := models.Login(reqData.Username, reqData.Password); ok {
			et := utils.EasyToken{}
			validation, err := et.ValidateToken(user.Token)
			if !validation {
				et = utils.EasyToken{
					Username: user.Username,
					Uid: int64(user.Id),
					Expires: time.Now().Unix() + 2 * 3600,
				}
				token, err = et.GetToken()
				fmt.Printf("Token = %s\n", token)
				if token == "" || err != nil {
					this.Data["json"] = errUserToken
					this.ServeJSON()
				} else {
					models.UpdateUserToken(user, token)
				}
			} else {
				token = user.Token
			}
			models.UpdateUserLastLogin(user)
			var returnData = &UserSuccessLoginData{user.Id,token, user.Username}
			this.Data["json"] = &Response{0, 0, "ok", returnData}
		} else {
			this.Data["json"] = &errNoUser
		}
	} else {
		this.Data["json"] = "no user or password"
	}
	this.ServeJSON()
}

func (this *UserController) GetLogin() {
	this.Ctx.WriteString("HEllo String")
}
