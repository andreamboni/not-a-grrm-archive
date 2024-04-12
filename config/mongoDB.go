package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDBCollection() (*mongo.Collection, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return nil, fmt.Errorf("error initializing mongodb: %v", err)
	}

	collection := client.Database("grrmarchive").Collection("blogposts")

	return collection, nil
}
