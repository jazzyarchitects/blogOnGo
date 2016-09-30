package main

import (
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}

func getSession() *mgo.Session{
	s,err := mgo.Dial("mongodb://localhost:27017")
	if err!=nil{
		panic(err)
	}
	return s
}