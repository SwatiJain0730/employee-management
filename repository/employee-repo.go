package repository

import (
	"employee-management-service/model"
)

//EmployeeRepo ...
type EmployeeRepo interface {
	CreateEmployee(employee model.Employee) (bool, int)
	UpdateEmployee(employee model.Employee) bool
	GetEmployee(name string) model.Employee
}
