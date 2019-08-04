package model

//EmployeeBase ...
type EmployeeBase struct {
	Name           string  `json:"employeeName" valid:"matches(^[a-zA-Z0-9-_]+$),required~employeeName cannot be empty"`
	ID             int     `json:"employeeId" valid:"required~employeeId cannot be empty"`
	DepartmentName string  `json:"departmentName" valid:"matches(^[a-zA-Z0-9-_]+$),required~departmentName cannot be empty"`
	Salary         float32 `json:"salary" valid:"required~salary cannot be empty"`
}

//Employee ...
type Employee struct {
	EmployeeBase
	Status string `json:"employeeStatus" valid:"required~Status cannot be empty"`
}
