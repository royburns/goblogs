package main

import (
	"github.com/astaxie/beego"
	"github.com/royburns/goblogs/controllers"
	"github.com/royburns/goblogs/controllers/admin"
)

func main() {
	//beego.PprofOn = true
	/*
		beego.RegisterController("/", &controllers.MainController{})
		beego.RegisterController("/reg", &controllers.RegController{})
		beego.RegisterController("/login", &controllers.LoginController{})
		beego.RegisterController("/:id([0-9]+)", &controllers.MainController{})
		beego.RegisterController("/admin/index", &admin.IndexController{})
		beego.RegisterController("/admin/login", &admin.LoginController{})
		beego.RegisterController("/admin/addblog", &admin.AddBlogController{})
		beego.RegisterController("/admin/editblog/:id([0-9]+)", &admin.EditBlogController{})
		beego.RegisterController("/admin/delblog/:id([0-9]+)", &admin.DelBlogController{})
		//*/
	beego.Router("/", &controllers.MainController{})
	beego.Router("/reg", &controllers.RegController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/:id([0-9]+)", &controllers.MainController{})
	beego.Router("/admin/index", &admin.IndexController{})
	beego.Router("/admin/login", &admin.LoginController{})
	beego.Router("/admin/addblog", &admin.AddBlogController{})
	beego.Router("/admin/editblog/:id([0-9]+)", &admin.EditBlogController{})
	beego.Router("/admin/delblog/:id([0-9]+)", &admin.DelBlogController{})
	beego.Run()
}
