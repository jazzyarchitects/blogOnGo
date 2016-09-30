package models


import (
	"encoding/json"
	"log"
	"io/ioutil"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int `json:"age"`
	Email     string `json:"email"`
}


func (u *User) Save() {
	filename := "./data/" + u.FirstName + u.LastName + ".json"
	data, err := json.Marshal(u)
	if err != nil {
		log.Println("Error saving user")
		return
	}
	log.Println("Data: " + string(data[:]))
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal("Error writing: " + err.Error())
	}
}