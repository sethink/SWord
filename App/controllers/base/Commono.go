package base

import (
	"SWord/App"
	"SWord/App/models"
	"SWord/App/tools/Helper"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type CommonController struct {
	beego.Controller
}

func (c *CommonController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		if username == "" || password == "" {
			c.Data["json"] = Helper.Write(1002, "", "参数有误")
			c.ServeJSON()
			c.StopRun()
		}

		var user models.User

		if err := models.Db.Where("`username`=? AND `password`=?", username, Helper.UsersPasswordEncrypt(password)).First(&user).Error; gorm.IsRecordNotFoundError(err) {
			c.Data["json"] = Helper.Write(1002, "", "账号或密码不正确")
			c.ServeJSON()
			c.StopRun()
		}

		sess, _ := App.Sessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		_ = sess.Set("id", user.Id)

		c.Data["json"] = Helper.Write(200, "", "登录成功")
		c.ServeJSON()
		c.StopRun()
	}
	c.TplName = "base/common/login.html"
}
