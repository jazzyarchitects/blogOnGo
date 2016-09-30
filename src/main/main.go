package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"blogOnGo/src/models"
)


func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/",apiHandler)


	http.Handle("/", r)

	log.Println("Listening on PORT 3000")
	http.ListenAndServe(":3000", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Listened to api"))
	u := models.User{FirstName: "Jibin", LastName: "Mathews2"}
	u.Save()
}