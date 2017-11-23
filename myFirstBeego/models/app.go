package models

import (
	"fmt"
    "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {

Id      int  `orm:"column(id);auto"`
User_name string
Age      int
Email   string
   Staff_id     string
    Position         string
    Extension_number string
    Telephone_number string
    Office_location  string

}
func (t *User) TableName() string {
	return "user"
}
func RegisterDB() {

    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:+ZkGR*GDx5zP@tcp(127.0.0.1:3306)/sallenBeego")
    orm.RunSyncdb("default", true, false)
}

func init () {
    fmt.Println("create table")
    orm.RegisterModel(new(User))
}

func UserAdd(o orm.Ormer,user User) (int64,error){

return o.Insert(&user)

}

func UserQuery(o orm.Ormer,sql string,paras [] interface{}) ([] User,error){
    // query
    var users [] User
    num,err := o.Raw(sql,paras).QueryRows(&users)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)
    return users,err
}