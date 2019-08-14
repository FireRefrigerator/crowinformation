package routers

import (
	"service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/v1/crowNews", &controllers.CrowNewsControllers{}, "*:CrowNews")
}
