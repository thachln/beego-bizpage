package routers

import (
	"beego-bizpage/controllers"
	beego "github.com/astaxie/beego/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
