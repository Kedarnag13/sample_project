package models

import (
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id                   int
	Username             string `form:"username"`
	Password             string `form:"password"`
	PasswordConfirmation string `form:"password_confirmation"`
}

func init() {
	orm.RegisterModel(new(Users))
}
