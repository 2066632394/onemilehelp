package routers

import (
	"szhapp/controllers"
	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
    admin.Run()
    beego.Router("/", &controllers.MainController{})
    beego.Router("/onemile/gethelps", &controllers.HelpController{}, "Get:GetHelps")
}
