package home

import (
	. "cms/controllers/rbac"
)

type HomeController struct {
	CommonController
}

func (c *HomeController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = c.GetTemplatetype() + "/home/index.tpl"
}
