package main

import (
	"Pureblog/controllers"
	_ "Pureblog/models"
	_ "Pureblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.Async(1e3)
	logs.SetLogger(logs.AdapterFile, `{"filename":"log/Pureblog.log"}`)
	beego.BConfig.EnableGzip = true
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
