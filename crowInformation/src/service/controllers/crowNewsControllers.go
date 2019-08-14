package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"service/model"
)

type CrowNewsControllers struct {
	beego.Controller
}

func (c *CrowNewsControllers) CrowNews() {
	fmt.Println("println hello")
	model.ProductNewsUrl(mainPagestr)
	model.ConsumeNewsUrl()
	c.Ctx.WriteString("hello crow news")
}

var mainPagestr = `<!doctype html>`
