package main

import (
	"fmt"
	"log"

	"github.com/golang-mongodb/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var collection *mgo.Collection

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("Could not connect to mongo: ", err.Error())
	}
	defer session.Close()
	if err != nil {
		fmt.Print(err)
	}
	defer session.Close()
	collection = session.DB("test").C("employee")
	if collection == nil {
		log.Fatal("Could not get collection employee")
		return
	}

	insertEmployee(getMockModel())

	// employeies := getAllEmployee()
	// if len(employeies) == 0 {
	// 	fmt.Println("Not have any Employee")
	// } else {
	// 	for _, item := range employeies {
	// 		fmt.Println(item.ID)
	// 	}
	// }

	// deleteEmployee(employeies[len(employeies)-1].ID)

	// employee := getEmployeeByID(employeies[len(employeies)-1].ID)
	// if employee == nil {
	// 	log.Println("Could not get employee")
	// } else {
	// 	log.Println(employee)
	// }
}

func getEmployeeByID(employeeID bson.ObjectId) (employee *model.Employee) {
	collection.Find(bson.M{"_id": employeeID}).One(&employee)
	return employee
}

func getAllEmployee() (employies []model.Employee) {
	collection.Find(bson.M{}).All(&employies)
	return employies
}

func insertEmployee(employeeModel model.Employee) {
	err := collection.Insert(employeeModel)
	if err != nil {
		log.Println("Error creating Employee: ", err.Error())
	} else {
		log.Println("Created Employee success")
	}
}

func deleteEmployee(employeeID bson.ObjectId) {
	err := collection.RemoveId(employeeID)
	if err != nil {
		log.Println("Error delete Employee: ", err.Error())
	} else {
		log.Println("Deleted Employee success")
	}
}

func getMockModel() (model model.Employee) {
	model.ID = bson.NewObjectId()
	model.Email = "golang@gmail.com"
	model.Password = "golang"
	model.Username = "Golang"
	return model
}
