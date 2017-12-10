package controllers

import (
	"github.com/astaxie/beego"
    "os/exec"
    "fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
    str, err := exec.Command("/bin/bash","./sn.sh").Output()
    if err != nil {
        fmt.Println(err.Error())
        c.Ctx.WriteString(err.Error())
        return
    }
    c.Ctx.WriteString(string(str))
}
