package models

import "gopkg.in/mgo.v2/bson"

type Blog struct {
	Id bson.ObjectId `bson:"_id" json:"_id"`
	Token string `bson:"token" json:"id"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	LikedBy []bson.ObjectId `bson:"likedBy" json:"likedBy"`
}
