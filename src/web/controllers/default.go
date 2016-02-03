package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
	c.Data["Website"] = "www.jaxiu.cn"
//	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Email"] = "axiu@jaxiu.cn"
	c.TplName = "index.tpl"

}
