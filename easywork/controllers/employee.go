package controllers

import (
	//"encoding/json"
	//"strconv"

	"github.com/astaxie/beego"
	"github.com/luck7/easywork/models"
)

type EmployeeController struct {
	beego.Controller
}

// Example:
//
//   req: GET /employee/
//   res: 200 {"Employees": [
//          {"ID": 1, "Title": "Learn Go", "Done": false},
//          {"ID": 2, "Title": "Buy bread", "Done": true}
//        ]}
func (this *EmployeeController) ListEmployees() {
	res := struct{ Employees []*models.Employee }{models.DefaultEmployeeList.All()}
	this.Data["json"] = res
	this.ServeJson()
}
