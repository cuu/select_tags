package routers

import (
	"github.com/cuu/select_tags/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include( &controllers.MainController{})
	beego.Include( &controllers.PtsController{})

	beego.Include( &controllers.ModelsGetSearchController{})
	
	
}
