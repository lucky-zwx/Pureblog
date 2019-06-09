package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.TplName = "404.html"
}

func (c *ErrorController) Error502() {
	c.TplName = "502.html"
}


func (c *ErrorController) ErrorDb() {
	c.TplName = "502.html"
}
