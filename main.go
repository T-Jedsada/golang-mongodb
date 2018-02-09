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

	employee := getEmployeeByID("9")
	if employee == nil {
		fmt.Println("Could not get employee")
	}
	fmt.Println(employee)

	employeies := getAllEmployee(session)
	if len(employeies) == 0 {
		fmt.Println("Not have any Employee")
	} else {
		for index, item := range employeies {
			fmt.Printf("index at %d : %s\n", index, item)
		}
	}
}

func getEmployeeByID(employeeID string) (employee *model.Employee) {
	collection.Find(bson.M{"id": employeeID}).One(&employee)
	return employee
}

func getAllEmployee(s *mgo.Session) (employies []model.Employee) {
	collection.Find(bson.M{}).All(&employies)
	return employies
}

func deleteEmployee(s *mgo.Session, employeeID string) (isSuccess bool) {
	return false
}

func addEmployee(s *mgo.Session, employee model.Employee) (isSuccess bool) {
	return false
}
