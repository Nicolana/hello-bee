package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func init() {
	// new(User) 表示实例化一个对象
	orm.RegisterModel(new(User))
}

type User struct {
	Id int `json:"id" orm:"column(id);pk;unique;auto_increment;int(11)"`
	Username string `json:"username" orm:"column(username);unique;size(32)"`
	Password string `json:"password" orm:"column(password);size(128)"`
	Avatar string `json:"avatar" orm:"column(avatar);varbinary"`
	Salt string `json:"salt" orm:"column(salt);size(128)"`
	Token string `json:"token" orm:"column(token);size(256)"`
	LastLogin int64 `json:"last_login" orm:"column(last_login);size(11)"`
	Status int `json:"status" orm:"column(status);size(1)"`
	CreatedAt int64 `json:"created_at" orm:"column(created_at);size(11)"`
	UpdatedAt int64 `json:"updated_at" orm:"column(updated_at);size(11)"`
}

func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetUserByUserName(username string) (v *User, err error) {
	o := orm.NewOrm()
	fmt.Println(username)
	v = &User{Username: username}
	if err := o.QueryTable(new(User)).Filter("Username", username).RelatedSel().One(v); err == nil {
		return  v, nil
	}
	return nil, err
}

func GetUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}

func Login(username string, password string) (bool, *User) {
	o := orm.NewOrm()
	fmt.Printf("username = %s\n", username)
	user, err := GetUserByUserName(username)
	if err != nil {
		return false, nil
	}
	err = o.QueryTable(user).Filter("Username", username).Filter("Password", password).One(user)
	return err != orm.ErrNoRows, user
}

func UpdateUserToken(m *User, token string) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	m.Token = token
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return err
}

func UpdateUserLastLogin(m *User) (err error){
	o := orm.NewOrm()
	v := User{Id: m.Id}
	lastLogin := time.Now().UTC().Unix()
	m.LastLogin = lastLogin
	if err := o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of record updated in atabase: ", num)
		}
	}
	return  err
}