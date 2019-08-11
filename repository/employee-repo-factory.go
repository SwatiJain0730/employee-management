package repository

import (
	"employee-management-service/database"
)

//CreateEmployeeDAO ...
func CreateEmployeeDAO() EmployeeRepo {
	dao := new(EmployeDBDAO)
	dao.mongoDbCollection = database.MongoDBCollection
	return dao
}
