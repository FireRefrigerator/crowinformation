package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type CrowNewsControllers struct {
	beego.Controller
}

func (c *CrowNewsControllers) CrowNews() {
	fmt.Println("println hello")
	c.Ctx.WriteString("hello crow news")
}
