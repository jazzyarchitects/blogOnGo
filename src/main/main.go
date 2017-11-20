package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"blogOnGo/src/controllers"
	"gopkg.in/mgo.v2"
	"fmt"
	"net/http/httptest"
	"net/http/httputil"
)


func logHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			x, err := httputil.DumpRequest(r, true)
			if err != nil {
					http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
					return
			}
			log.Println(fmt.Sprintf("%q", x))
			rec := httptest.NewRecorder()
			fn(rec, r)
			log.Println(fmt.Sprintf("%q", rec.Body))            
	}
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "A message was received")
}

func main() {
	router := NewRouter()
	router.HandleFunc("/", logHandler(MessageHandler))
	log.Fatal(http.ListenAndServe(":3005", router))
}

func NewRouter() *mux.Router {
	r := mux.NewRouter();
	router := r.PathPrefix("/api").Subrouter()

	session,err := mgo.Dial("mongodb://localhost:27017")
	if err!=nil{
		panic(err)
	}

	fmt.Println("Listening on port 3005...");

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

	//----------------Ping Routes-------------------
	pingRouter := router.PathPrefix("/ping").Subrouter()
	pingController := controllers.NewPingController()

	pingRouter.HandleFunc("/", pingController.RespondPing).Methods("GET")
	pingRouter.HandleFunc("/swagger", pingController.SendSwagger).Methods("GET")
	

	return router
}
