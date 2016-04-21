package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Id    bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Phone string        `bson:"phone"`
}
