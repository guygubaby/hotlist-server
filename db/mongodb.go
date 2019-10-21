package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func InitMongo() (*mongo.Collection,context.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://guygubaby.top:27017"))

	if err!=nil {
		log.Fatal(err)
	}

	collection := client.Database("hotlist").Collection("list")
	return collection,ctx
}
