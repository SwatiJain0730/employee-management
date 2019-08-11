package service

import (
	"employee-management-service/repository"
)

func init() {
	employeeRepo = repository.CreateEmployeeDAO
}
