package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "myFirstBeego/models"
    "fmt"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {
o := orm.NewOrm()

    user := models.User{
                User_name:"loongshawn",
                Age:29,
                Email:"loongshawn@jjj.com",
                Staff_id:"123456",
                Position:"研发工程师",
                Extension_number:"",
                Telephone_number:"15642314234",
                Office_location:"商会产业园",
                }
  o.Begin()
    id, err := models.UserAdd(o,user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)
    if err != nil {
        o.Rollback()
        fmt.Println("插入user表出错,事务回滚")
    } else {
        o.Commit()
        fmt.Println("插入user表成功,事务提交")
    }
    c.TplName = "hello-sitepoint.tpl"


}