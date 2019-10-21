package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type HotItem struct {
	ID primitive.ObjectID  `json:"_id" bson:"_id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Desc *string `json:"desc"`
	Cate int `json:"cate"`
}
