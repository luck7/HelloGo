package controllers

import (
	"github.com/astaxie/beego"
	. "github.com/luck7/easywork/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["Website"] = "luck7.github.io"
	c.Data["Email"] = "luck7@live.cn"
	c.TplNames = "user/index.tpl"

	var us User
	c.Data["JsonString"] = us.GetString()
}
