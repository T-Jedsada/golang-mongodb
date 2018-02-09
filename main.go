package main

import (
	"fmt"
	"log"

	"github.com/golang-mongodb/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("Could not connect to mongo: ", err.Error())
	}
	defer session.Close()
	if err != nil {
		fmt.Print(err)
	}

	employee := getEmployeeByID(session, "9")
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

func getEmployeeByID(s *mgo.Session, employeeID string) (employee *model.Employee) {
	session := s.Copy()
	defer session.Close()
	c := session.DB("test").C("employee")
	c.Find(bson.M{"id": employeeID}).One(&employee)
	return employee
}

func getAllEmployee(s *mgo.Session) (employies []model.Employee) {
	session := s.Copy()
	defer session.Close()
	c := session.DB("test").C("employee")
	c.Find(bson.M{}).All(&employies)
	return employies
}

func deleteEmployee(s *mgo.Session, employeeID string) bool {
	return false
}

func addEmployee(s *mgo.Session, employee model.Employee) {

}
