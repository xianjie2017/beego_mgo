package routers

import (
	"github.com/astaxie/beego"
	"starchild/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello/get/:name", &controllers.HelloController{}, "get:Get")
	beego.Router("/hello/delete/:name", &controllers.HelloController{}, "get:Delete")
	beego.Router("/hello/update/:name/:phone", &controllers.HelloController{}, "get:Put")
	beego.Router("/hello/:name/:phone", &controllers.HelloController{}, "get:Post")
	beego.Router("/hello", &controllers.HelloController{})
}
