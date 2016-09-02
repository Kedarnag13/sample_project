package routers

import (
	"github.com/Kedarnag13/sample_project/controllers"
	"github.com/Kedarnag13/sample_project/controllers/users"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("/users/index", &users.UsersController{}, "get:Index")
	beego.Router("/users/new", &users.UsersController{}, "get:New")
	beego.Router("/users/sign_up", &users.UsersController{}, "post:Create")
}
