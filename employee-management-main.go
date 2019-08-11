package main

import (
	"employee-management-service/controller"
	"employee-management-service/database"
)

func main() {

	database.CreateDB()

	router := controller.RegisterRoutes()
	//go router.Run()
	router.Run()
}
