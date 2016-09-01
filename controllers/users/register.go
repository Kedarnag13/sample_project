package users

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UsersController struct {
	beego.Controller
}

type User struct {
	Username             string `form:"username"`
	Password             string `form:"password"`
	PasswordConfirmation string `form:"password_confirmation"`
}

func (u *UsersController) New() {
	u.TplName = "users/new.tpl"
}

func (u *UsersController) Post() {
	user := User{}
	if err := u.ParseForm(&user); err != nil {
		panic(err)
	}
	fmt.Printf("Username: %v", user.Username)
	u.TplName = "users/post.tpl"
}
