package main

import (
	"fmt"
	"log"

	"github.com/golang-mongodb/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	db, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	defer db.Close()
	c := db.DB("test").C("employee")
	if err != nil {
		fmt.Print(err)
	}

	var employee []model.Employee

	c.Find(bson.M{}).All(&employee)

	for index, item := range employee {
		fmt.Printf("index at %d : %s\n", index, item)
	}
}
