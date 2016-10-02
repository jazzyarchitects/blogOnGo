package models

import (
	_ "encoding/json"
	_ "log"
	_ "io/ioutil"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	FirstName string `bson:"fname" json:"fname"`
	LastName  string `bson:"lname" json:"lname"`
	Age       uint64 `bson:"age" json:"age"`
	Email     string `bson:"email" json:"email"`
}

type Users []User
