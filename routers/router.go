package routers

import (
	"github.com/Kedarnag13/sample_project/controllers/users"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &users.UsersController{})
}
