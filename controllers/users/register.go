package users

import (
	"github.com/Kedarnag13/sample_project/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type UsersController struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=password dbname=sample_project_development sslmode=disable")
	orm.RunCommand()
}

func (u *UsersController) New() {
	u.TplName = "users/new.tpl"
}

func (u *UsersController) Post() {

	o := orm.NewOrm()
	o.Using("default")

	user := models.Users{}
	if err := u.ParseForm(&user); err != nil {
		panic(err)
	}

	us := new(models.Users)
	us.Username = user.Username
	us.Password = user.Password
	us.PasswordConfirmation = user.PasswordConfirmation

	o.Insert(us)

	u.Ctx.Redirect(302, "/users/new")
	// fmt.Printf("Username: %v", user.Username)
	// u.TplName = "users/post.tpl"
}
