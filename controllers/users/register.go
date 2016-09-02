package users

import (
	"fmt"
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

func (u *UsersController) Index() {
	o := orm.NewOrm()
	var lists []orm.Params
	num, err := o.QueryTable("users").Values(&lists)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, row := range lists {
			for range row {
				fmt.Println(row)
				u.Data["user"] = row
			}
		}
	}
	u.TplName = "users/index.tpl"
}

func (u *UsersController) New() {
	u.TplName = "users/new.tpl"
}

func (u *UsersController) Create() {

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

	success, err := o.Insert(us)

	if err != nil || success == 0 {
		u.Ctx.Redirect(302, "/users/new")
	} else {
		u.Ctx.Redirect(302, "/users/index")
	}

}
