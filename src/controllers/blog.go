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

/*
This will retrieve a list of all the blogs stored in the database\
 */
func (bc BlogController)GetBlogFeed(w http.ResponseWriter, r *http.Request) {
	pageQuery := r.URL.Query()["page"]

	var page int;

	if len(pageQuery) > 0 {
		page, _ = strconv.Atoi(pageQuery[0])
	} else {
		page = 1
	}
	var blogs []models.Blog

	err := blogCollection.Find(nil).Limit(blogLimit).Skip((page - 1) * blogLimit).All(&blogs)
	if err != nil {
		w.WriteHeader(404)
	}

	blogsJson, err := json.Marshal(blogs)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", blogsJson)
}


/*
This will retrieve a blog stored in the database by id or Token
 */
func (bc BlogController)GetBlogByToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	id := p["id"]

	fmt.Println(id)

	var blog models.Blog;

	err := blogCollection.Find(bson.M{"token": id}).One(&blog)

	if err != nil && err.Error() == "not found" {
		w.WriteHeader(200)
		fmt.Fprint(w, "{\"success\": true, \"data\": []}")
		return
	}
	//fmt.Println(err)
	fmt.Println(blog)

	blogJSON, _ := json.Marshal(blog)

	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", blogJSON)
}


/*
Create a new blog in the database. This function generates a unique token to identify the blog using only 7 characters
 */
func (bc BlogController)CreateBlog(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	blog := models.Blog{
		Title: r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	blog.Id = bson.NewObjectId()

	blog.Token = getUniqueToken(blogCollection)

	blogCollection.Insert(blog)

	b, _ := json.Marshal(blog)

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", b)

}

func getUniqueToken(collection *mgo.Collection) string {
	token := getRandomString(7)
	var blogs []models.Blog
	collection.Find(bson.M{token: token}).All(&blogs)
	if len(blogs) == 0 {
		return token
	} else {
		return getUniqueToken(collection)
	}
}

func (bc BlogController)UpdateBlog(w http.ResponseWriter, r *http.Request) {
	//p := mux.Vars(r)
	//id := p["id"]

}

