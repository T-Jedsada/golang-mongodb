package model

import "gopkg.in/mgo.v2/bson"

// Employee is model for mapping from db
type Employee struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Email    string        `json:"email"`
}
