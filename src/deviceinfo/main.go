package main

import (
	_ "deviceinfo/routers"
	"github.com/astaxie/beego"
)

func main() {
    beego.SetStaticPath("/swagger", "swagger")
	beego.Run()
}

