package service

import (
	"employee-management-service/model"
	"employee-management-service/repository"
	"fmt"
)

var employeeRepo func() repository.EmployeeRepo

//CreateEmployee ...
func CreateEmployee(employee model.Employee) (bool, int) {
	ok, dbID := employeeRepo().CreateEmployee(employee)
	return ok, dbID
}

//GetEmployee ...
func GetEmployee(name string) model.Employee {
	fmt.Println("GetEmployee", name)
	employee := employeeRepo().GetEmployee(name)
	return employee
}

//UpdateEmployee ...
func UpdateEmployee(employee model.Employee) bool {
	result := employeeRepo().UpdateEmployee(employee)
	return result
}
