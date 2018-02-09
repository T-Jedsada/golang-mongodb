package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
)

func main() {
	db, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	defer db.Close()
	c := db.DB("test").C("employee")
	count, err := c.Count()
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%d\n", count)
	}
}
