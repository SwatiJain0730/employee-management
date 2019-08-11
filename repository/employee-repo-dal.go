package repository

import (
	"context"
	"employee-management-service/model"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//EmployeDBDAO ...
type EmployeDBDAO struct {
	mongoDbCollection *mongo.Collection
}

//CreateEmployee ...
func (collection *EmployeDBDAO) CreateEmployee(employee model.Employee) (bool, int) {
	fmt.Println("here")
	insertResult, err := collection.mongoDbCollection.InsertOne(context.TODO(), employee)
	if err != nil {
		log.Fatal(err)
		return false, 0
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	fmt.Println(insertResult)
	return true, 1
}

//GetEmployee ...
func (collection *EmployeDBDAO) GetEmployee(name string) model.Employee {
	var employee model.Employee
	filter := bson.D{{"name", name}}
	fmt.Println(name)
	err := collection.mongoDbCollection.FindOne(context.TODO(), filter).Decode(&employee)
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err)
		return model.Employee{}
	}
	fmt.Println("Found a single document: ", employee)
	return employee
}

//UpdateEmployee ...
func (collection *EmployeDBDAO) UpdateEmployee(employee model.Employee) bool {
	fmt.Println("document to update ", employee)
	filter := bson.D{{"name", employee.Name}}
	update := bson.D{
		{"$set", bson.D{
			{"salary", employee.Salary},
		}},
	}
	fmt.Println("Filter ", filter)
	fmt.Println("Update ", update)
	updateResult, err := collection.mongoDbCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		//return false, 0
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return true
}
