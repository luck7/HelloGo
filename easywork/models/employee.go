package models

import (
	//"encoding/json"
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

var DefaultEmployeeList *EmployeeManager

// EmployeeManager manages a list of employees in memory.
type EmployeeManager struct {
	employees []*Employee
	lastID    int64
}

type Employee struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(100)"`
	Dept string `orm:"size(100)"`
	Tele string `orm:"size(100)"`
	Mobile string `orm:"size(100)"`
	Email string `orm:"size(100)"`
	Place string `orm:"size(100)"`
	Ip string `orm:"size(100)"`
}

type Employeeslice struct {
	Employees []Employee
}

func init() {
	// register model
	orm.RegisterModel(new(Employee))

	// set default database
	orm.RegisterDataBase("default", "sqlite3", "easywork.sqlite", 30)

	DefaultEmployeeList = NewEmployeeManager()
}

// NewEmployeeManager returns an empty EmployeeManager.
func NewEmployeeManager() *EmployeeManager {
	return &EmployeeManager{}
}

// All returns the list of all the Tasks in the TaskManager.
func (m *EmployeeManager) All() []*Employee {
	o := orm.NewOrm()

	var emps []*Employee
	qs := o.QueryTable("employee")
	num, err := qs.All(&emps)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Get Employees:", num)	
	return emps
}
