package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	maxIdle := 30
	maxConn := 100
	mysqlurls := beego.AppConfig.String("mysqlurls")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", ""+mysqlurls, maxIdle, maxConn)
	orm.RegisterModel(new(BlogArticle), new(Admin))
	orm.RunSyncdb("default", false, true)
}

type BlogArticle struct {
	Id          int64 `orm:"auto;index;pk"`
	Author      string
	Top         bool
	Title       string
	Content     orm.TextField
	Morecontent orm.TextField
	Category    string
	Addtime     orm.DateTimeField `orm:"null"`
}

type Admin struct {
	Username string `orm:"pk"`
	Pappwd   string
	Headimg  string
	Email    string
}
