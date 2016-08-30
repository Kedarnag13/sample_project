package users

import (
	"github.com/astaxie/beego"
)

type UsersController struct {
	beego.Controller
}

func (c *UsersController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "users/index.tpl"
}
