package controllers

import (
	"gopkg.in/mgo.v2"
	"net/http"
	"blogOnGo/src/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type BlogController struct {
	session *mgo.Session
}

var blogCollection *mgo.Collection
var blogLimit int = 10

const BlogCollectionName string = "Blog"

func NewBlogController(s *mgo.Session) *BlogController {
	blogCollection = s.DB(DatabaseName).C(BlogCollectionName)
	return &BlogController{s}
}

func (bc BlogController)GetBlogFeed(w http.ResponseWriter, r *http.Request) {
	pageQuery := r.URL.Query()["page"]

	var page int;

	if len(pageQuery) > 0 {
		page,_ = strconv.Atoi(pageQuery[0])
	}else{
		page = 1
	}
	var blogs []models.Blog

	err := blogCollection.Find(nil).Limit(blogLimit).Skip((page-1)*blogLimit).All(&blogs)
	if err!=nil{
		w.WriteHeader(404)
	}

	blogsJson, err := json.Marshal(blogs)

	if err!=nil{
		panic(err)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", blogsJson)
}

func (bc BlogController)GetBlogByToken(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	id := p["id"]

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	//oid := bson.ObjectIdHex(id)

	var blog models.Blog;
	err := blogCollection.Find(bson.M{}).One(&blog)

	if err != nil {
		w.WriteHeader(500)
		panic(err)
	}

	blogJSON, _ := json.Marshal(blog)

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", blogJSON)
}

func (bc BlogController)CreateBlog(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
}

func (bc BlogController)UpdateBlog(w http.ResponseWriter, r *http.Request) {
	//p := mux.Vars(r)
	//id := p["id"]

}

