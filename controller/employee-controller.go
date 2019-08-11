package controller

import (
	"employee-management-service/model"
	"employee-management-service/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	funk "github.com/thoas/go-funk"
)

//RegisterRoutes registers routes
func RegisterRoutes() *gin.Engine {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	store := sessions.NewCookieStore([]byte("sessionSuperSecret"))
	router.Use(sessions.Sessions("sessionName", store))

	empRoutes := router.Group("/v1/api")

	empRoutes.POST("/login", loginHandler)
	empRoutes.GET("/logout", logoutHandler)

	empRoutes.GET("/employees/:name", AuthenticationRequired(), GetEmployee)
	empRoutes.POST("/employees", AuthenticationRequired(), CreateEmployee)
	empRoutes.PUT("/employees", AuthenticationRequired(), UpdateEmployee)
	return router
}

var (
	VALID_AUTHENTICATIONS = []string{"user", "admin", "subscriber"}
)

func loginHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)

	if strings.Trim(user.Username, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username can't be empty"})
	}
	if !funk.ContainsString(VALID_AUTHENTICATIONS, user.AuthType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid auth type"})
	}

	// Note: This is just an example, in real world AuthType would be set by business logic and not the user
	session.Set("user", user.Username)
	session.Set("authType", user.AuthType)

	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate session token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "authentication successful"})
}

func logoutHandler(c *gin.Context) {
	session := sessions.Default(c)

	// this would only be hit if the user was authenticated
	session.Delete("user")
	session.Delete("authType")

	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate session token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged out"})

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

	fmt.Println(c.Params)
	//employeeName := "Swati"
	employeeName := c.Params.ByName("name")
	employee := service.GetEmployee(employeeName)
	c.JSON(http.StatusOK, gin.H{"success": true, "employee": employee})
	return
}

//UpdateEmployee ...
func UpdateEmployee(c *gin.Context) {
	fmt.Println("Updating Employee")

	var createEmployeeStruct struct {
		model.EmployeeBase
		Status string `json:"employeeStatus"`
	}
	err := c.BindJSON(&createEmployeeStruct)
	jsonPayload, _ := json.Marshal(createEmployeeStruct)
	fmt.Println("Employee Update payload json", string(jsonPayload))
	if err != nil {
		errMessage := fmt.Sprintf("Cannot update empoyee becasue of error : [%v]", err)
		fmt.Println(errMessage)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": errMessage})
		return
	}

	employee := model.Employee{
		EmployeeBase: createEmployeeStruct.EmployeeBase,
		Status:       "ACTIVE",
	}
	updated := service.UpdateEmployee(employee)
	if updated {
		employeeUpdateMsg := fmt.Sprintf("Eployee data updated successfully with employee id [%b]", updated)
		fmt.Println(employeeUpdateMsg)
		c.JSON(http.StatusOK, gin.H{"success": true, "message": employeeUpdateMsg})
	} else {
		fmt.Println("Employee updation failed")
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Employee updation failed"})
	}
}
