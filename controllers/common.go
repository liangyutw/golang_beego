package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type Common struct {
	web.Controller
}

func (this *Common) Prepare() {
	this.Layout = "homelayout.html"
}
