package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
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
	Id          int64       `orm:"column(id);auto"`
	Author      string    `orm:"column(author);size(255)"`
	Top         bool      `orm:"column(top)"`
	Title       string    `orm:"column(title);size(255)"`
	Content     orm.TextField    `orm:"column(content)"`
	Morecontent orm.TextField    `orm:"column(morecontent)"`
	Category    string    `orm:"column(category);size(255)"`
	Addtime     time.Time `orm:"column(addtime);type(datetime);null;auto_now_add"`
}

type Admin struct {
	Username      string    `orm:"column(username);pk"`
	Pappwd  string `orm:"column(pappwd);size(255)"`
	Headimg string `orm:"column(headimg);size(255)"`
	Email   string `orm:"column(email);size(255)"`
}

type Session struct {
	Id            int    `orm:"column(session_key);pk"`
	SessionData   string `orm:"column(session_data);null"`
	SessionExpiry uint   `orm:"column(session_expiry)"`
}
