package controller

import (
	"github.com/guygubaby/hotlist-server/db"
	"github.com/guygubaby/hotlist-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func GetList(cate int) []model.HotItem {
	var res []model.HotItem
	collection,ctx := db.InitMongo()
	cur,_ := collection.Find(ctx,bson.M{"cate":cate})
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result model.HotItem
		err := cur.Decode(&result)
		if err !=nil {
			log.Fatal("Decode error : ",err)
		}
		res = append(res, result)
	}
	return res
}

