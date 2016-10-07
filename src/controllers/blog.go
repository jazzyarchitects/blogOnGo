package controllers

import (
	"gopkg.in/mgo.v2"
	"net/http"
	"blogOnGo/src/models"
	"encoding/json"
	"fmt"
)

type BlogController struct {
	session *mgo.Session
}

var blogCollection *mgo.Collection
var blogLimit int = 10

const BlogCollectionName string = "Blog"

func NewBlogController(s *mgo.Session) *BlogController{
	blogCollection = s.DB(DatabaseName).C(BlogCollectionName)
	return &BlogController{s}
}

func (bc BlogController)getBlogFeed(w http.ResponseWriter, r *http.Request){
	page := r.URL.Query()["page"]
	if page==nil{
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
	fmt.Fprintf(w, "%s", blogsJson)
}

func 
