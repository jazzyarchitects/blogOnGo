package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"blogOnGo/src/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}

func NewRouter() *mux.Router {
	r := mux.NewRouter();
	router := r.PathPrefix("/api").Subrouter()

	session,err := mgo.Dial("mongodb://localhost:27017")
	if err!=nil{
		panic(err)
	}

	//----------------User Routes-------------------

	userRouter := router.PathPrefix("/user").Subrouter()
	userController := controllers.NewUserController(session)

	userRouter.HandleFunc("/{id}", userController.GetUserById).Methods("GET")
	userRouter.HandleFunc("/{id}", userController.RemoveUser).Methods("DELETE")
	userRouter.HandleFunc("/", userController.GetUsers).Methods("GET")
	userRouter.HandleFunc("/", userController.CreateUser).Methods("POST")

	//----------------Blog Routes-------------------

	blogRouter := router.PathPrefix("/blog").Subrouter()
	blogController := controllers.NewBlogController(session)

	blogRouter.HandleFunc("/{id}", blogController.GetBlogByToken).Methods("GET")
	blogRouter.HandleFunc("/{id}", blogController.UpdateBlog).Methods("PUT")
	blogRouter.HandleFunc("/", blogController.GetBlogFeed).Methods("GET")
	blogRouter.HandleFunc("/", blogController.CreateBlog).Methods("POST")

	return router
}
