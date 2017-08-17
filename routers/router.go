package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/", &controllers.MainController{})
  beego.Router("/bills/new", &controllers.BillController{},"get:New")
  beego.Router("/bills/create", &controllers.BillController{},"post:Create")
  beego.Router("/bills/:id", &controllers.BillController{},"get:Show")
  beego.Router("/bills", &controllers.BillController{},"get:Index")
  beego.Router("/bills/:id/edit", &controllers.BillController{},"get:Edit")
  beego.Router("/bills/:id/update", &controllers.BillController{},"post:Update")
  beego.Router("/bills/:id/print", &controllers.BillController{},"get:Print")
}
