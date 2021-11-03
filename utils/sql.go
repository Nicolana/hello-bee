package utils

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	dbuser, _ := beego.AppConfig.String("db.user")
	dbpassword, _ := beego.AppConfig.String("db.password")
	dbhost, _ := beego.AppConfig.String("db.host")
	dbport, _ := beego.AppConfig.String("db.port")
	dbname, _ := beego.AppConfig.String("db.name")
	dbcharset, _ := beego.AppConfig.String("db.charset")
	runmode, _ := beego.AppConfig.String("runmode")
	if runmode == "dev" {
		orm.Debug = true
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + dbcharset + "&net_write_timeout=6000"
	fmt.Println("连接到: " + dsn)
	orm.RegisterDataBase("default", "mysql", dsn)
	if err := orm.RunSyncdb("default", false, true); err != nil {
		fmt.Println(err)
	}
	orm.RunCommand()
}
