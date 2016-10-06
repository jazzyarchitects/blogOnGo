package models

import "gopkg.in/mgo.v2/bson"

type Blog struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Likes int64 `bson:"likes" json:"likes"`
	LikedBy []bson.ObjectId `bson:"likedBy" json:"likedBy"`
}
