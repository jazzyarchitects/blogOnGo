package models

import "gopkg.in/mgo.v2/bson"

type Blog struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	FirstName string `bson:"fname" json:"fname"`
	LastName  string `bson:"lname" json:"lname"`
	Age       uint64 `bson:"age" json:"age"`
	Email     string `bson:"email" json:"email"`
	Password 	string `bson:"password" json:"password"`
}
