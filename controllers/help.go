package controllers

import (
	"github.com/astaxie/beego"
)

type HelpController struct {
	beego.Controller
}


func (c *HelpController) Prepare() {

}

func (c *HelpController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
