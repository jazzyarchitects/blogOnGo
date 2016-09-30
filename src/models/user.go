package models

import (
	"encoding/json"
	"log"
	"io/ioutil"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID bson.ObjectId `bson:"_id",omitempty`
	FirstName string
	LastName  string
	Age       int
	Email     string
}

type Users []User

func (u *User) Save() {
	//filename := "./data/" + u.FirstName + u.LastName + ".json"
	//data, err := json.Marshal(u)
	//if err != nil {
	//	log.Println("Error saving user")
	//	return
	//}
	//log.Println("Data: " + string(data[:]))
	//err = ioutil.WriteFile(filename, data, 0644)
	//if err != nil {
	//	log.Fatal("Error writing: " + err.Error())
	//}
}

func (u *User)GetAllUsers() Users {

}