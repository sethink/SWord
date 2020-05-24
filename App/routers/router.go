package routers

import (
	"SWord/App"
	"SWord/App/controllers/base"
	"SWord/App/controllers/index"
	"SWord/App/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/jinzhu/gorm"
	"net/http"
)

func init() {
	beego.Get("/", func(c *context.Context) {
		c.Redirect(http.StatusFound, "/index/Index/index")
	})

	Base()
	Index()
	Admin()
}

func Base() {
	ns := beego.NewNamespace("/base",
		beego.NSAutoRouter(&base.CommonController{}),
	)
	beego.AddNamespace(ns)
}

func Index() {
	ns := beego.NewNamespace("/index",
		beego.NSBefore(IndexMiddleware),
		beego.NSAutoRouter(&index.IndexController{}),
	)
	beego.AddNamespace(ns)
}

func Admin() {

}

/******************************中间件 start************************************/
func IndexMiddleware(c *context.Context) {
	sess, _ := App.Sessions.SessionStart(c.ResponseWriter, c.Request)
	id := sess.Get("id")
	if id == nil {
		c.Redirect(http.StatusFound, "/base/Common/login")
		return
	}

	var user models.User
	if err := models.Db.Where("`id`=>", id).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		c.Redirect(http.StatusFound, "/base/Common/login")
		return
	}

	//if

	c.Input.SetData("userInfo", user)
}

/******************************中间件 end************************************/
