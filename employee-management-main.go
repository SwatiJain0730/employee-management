package main

import "employee-management-service/controller"

func main() {

	router := controller.RegisterRoutes()
	//go router.Run()
	router.Run()

}
