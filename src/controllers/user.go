package controllers

import (
	"gopkg.in/mgo.v2"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"blogOnGo/src/models"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"strconv"
	"encoding/json"
)

const DatabaseName string = "BlogOnGo"
const UserCollectionName string = "User"


type(
	UserController struct{
		session *mgo.Session
	}
)

var userCollection *mgo.Collection


func NewUserController(s *mgo.Session) *UserController{
	userCollection = s.DB(DatabaseName).C(UserCollectionName)
	return &UserController{s}
}


func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	var us []models.User

	//query := r.URL.Query()

	err := userCollection.Find(nil).All(&us)

	if err!=nil{
		w.WriteHeader(404)
	}

	log.Println(us)

	usersJson, err:= json.Marshal(us)

	if err!=nil{
		panic(err)
	}

	//log.Println(usersJson)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", usersJson)

}

func (uc UserController) GetUserById(w http.ResponseWriter, r *http.Request){
	p := mux.Vars(r)
	id := p["id"]

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)

	u := models.User{}

	err := userCollection.FindId(oid).One(&u)

	if err!=nil{
		panic(err)
		w.WriteHeader(404)
		return
	}

	uj,_ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

/*
 * Creates user from form values in post request and saves them in database
 */
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request){

	r.ParseForm()
	log.Println(r.Form)

	age,err := strconv.ParseUint(r.FormValue("age"),10,64)
	password := getPasswordHash(r.FormValue("password"))

	if err!=nil{
		panic(err)
	}
	u := models.User{
		FirstName: r.FormValue("fname"),
		LastName: r.FormValue("lname"),
		Age: age,
		Email: r.FormValue("email"),
		Password: password,
	}


	u.Id = bson.NewObjectId()

	userCollection.Insert(u)

	uj,err := json.Marshal(u)

	if err!=nil{
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf( w, "%s" ,uj)
	//w.Write([]byte(uj))
}

func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request){
	p := mux.Vars(r)
	id := p["id"]

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)

	err := userCollection.RemoveId(oid)

	if err!=nil{
		w.WriteHeader(404)
	}

	w.WriteHeader(200)
}