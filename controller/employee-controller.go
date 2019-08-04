package controller

import (
	"employee-management-service/model"
	"employee-management-service/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

//RegisterRoutes registers routes
func RegisterRoutes() *gin.Engine {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	empRoutes := router.Group("/v1")

	empRoutes.GET("/services", GetEmployee)
	empRoutes.POST("/services", CreateEmployee)
	return router
}

//CreateEmployee ...
func CreateEmployee(c *gin.Context) {
	fmt.Println("Creating Employee")

	var createEmployeeStruct struct {
		model.EmployeeBase
		Status string `json:"employeeStatus"`
	}
	err := c.BindJSON(&createEmployeeStruct)
	jsonPayload, _ := json.Marshal(createEmployeeStruct)
	fmt.Println("Employee payload json", string(jsonPayload))
	if err != nil {
		errMessage := fmt.Sprintf("Cannot create empoyee becasue of error : [%v]", err)
		fmt.Println(errMessage)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": errMessage})
		return
	}
	var employeePayloadErrors error
	valid, err := govalidator.ValidateStruct(createEmployeeStruct)
	if valid {
		if createEmployeeStruct.Status == "INACTIVE" {
			fmt.Println("Cannot create an inactive employee")
			employeePayloadErrors = errors.New("Cannot create an inactive employee")
		}
	} else {
		employeePayloadErrors = err
	}

	if employeePayloadErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Payload verification failed"})
		return
	}
	employee := model.Employee{
		EmployeeBase: createEmployeeStruct.EmployeeBase,
		Status:       "ACTIVE",
	}
	created, id := service.CreateEmployee(employee)
	if created {
		employeeCreateMsg := fmt.Sprintf("Eployee data created successfully with employee id [%s]", id)
		fmt.Println(employeeCreateMsg)
		c.JSON(http.StatusOK, gin.H{"success": true, "message": employeeCreateMsg})
	} else {
		fmt.Println("Employee creation failed")
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Employee creation failed"})
	}

}

//GetEmployee ...
func GetEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Hello"})
	return
}
