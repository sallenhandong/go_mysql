package routers

import (
	"myFirstBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello-world", &controllers.MainController{}, "get:HelloSitepoint")
	beego.Router("/login", &controllers.LoginController{}, "get:Login")
}
