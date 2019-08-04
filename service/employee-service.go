package service

import (
	"employee-management-service/model"
	"math/rand"
)

//CreateEmployee ...
func CreateEmployee(employee model.Employee) (bool, int) {
	min := 1
	max := 300
	return true, rand.Intn(max-min) + min
}
