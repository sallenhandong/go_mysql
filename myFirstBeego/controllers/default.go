package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// func (c *MainController) Get() {
// 	c.Data["Website"] = "beego.me"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.tpl"
// }
func (c *MainController) HelloSitepoint() {
    c.Data["Website"] = "My Website"
    c.Data["Email"] = "your.email.address@example.com"
    c.Data["EmailName"] = "Your Name"

    c.TplName = "hello-sitepoint.tpl"
}
