package main

import (
	_ "myFirstBeego/routers"
	"github.com/astaxie/beego"
	"myFirstBeego/models"
	
)

func init() {
 models.RegisterDB()
}

func main() {

	beego.Run()
}

