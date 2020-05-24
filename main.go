package main

import (
	_ "SWord/App"
	_ "SWord/App/models"
	_ "SWord/App/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "./assets/static")
	beego.SetViewsPath("App/views")

	beego.Run()
}
