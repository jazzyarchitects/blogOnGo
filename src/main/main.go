package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age int `json:"age"`
	Email string `json:"email"`
}

func (u *User) save() {
	filename := "./data/"+u.FirstName+u.LastName+".json"
	data, err := json.Marshal(u)
	if err!= nil{
		log.Println("Error saving user")
		return
	}
	log.Println("Data: "+string(data[:]))
	err = ioutil.WriteFile(filename, data, 0644)
	if err!=nil{
		log.Fatal("Error writing: "+err.Error())
	}
}


func main(){
	http.HandleFunc("/",mainHandler)


	log.Println("Listening on PORT 3000")
	http.ListenAndServe(":3000", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Heard "+r.URL.Path[1:]))
	u := User{FirstName: "Jibin", LastName: "Mathews"}
	u.save()
}
