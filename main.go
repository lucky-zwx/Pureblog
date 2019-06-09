package main

import (
	"Pureblog/controllers"
	_ "Pureblog/routers"
	"github.com/astaxie/beego"
	_ "Pureblog/models"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.Async(1e3)
	logs.SetLogger(logs.AdapterFile, `{"filename":"log/Pureblog.log"}`)
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

