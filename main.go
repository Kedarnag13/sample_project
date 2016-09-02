package main

import (
	"fmt"
	_ "github.com/Kedarnag13/sample_project/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {

	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	beego.Run()

}
