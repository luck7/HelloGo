package routers

import (
	"github.com/astaxie/beego"
	"github.com/luck7/easywork/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

//	beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
//	beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")
		
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/employee/", &controllers.EmployeeController{},"get:ListEmployees")
}
