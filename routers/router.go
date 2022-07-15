package routers

import (
	"demoProject/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/contactus", &controllers.ContactusController{})
	web.Router("/contactusInsert", &controllers.ContactusController{}, "POST:InsertUser")
	web.Router("/contactusThankyou", &controllers.ContactusController{}, "GET:Thankyou")
	web.Router("/contactusUpdate", &controllers.ContactusController{}, "POST:UpdateUser")
	web.Router("/contactusDel/:delId", &controllers.ContactusController{}, "GET:DelUser")
}
