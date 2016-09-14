package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id                   int
	Username             string `form:"username"`
	Password             string `form:"password"`
	PasswordConfirmation string `form:"password_confirmation"`
}

func init() {
	orm.RegisterModel(new(User))
}

func FetchAllUsers() ([]*User, error) {
	var users []*User
	_, err := orm.NewOrm().QueryTable(new(User)).All(&users)
	return users, err
}
