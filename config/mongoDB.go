package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDBCollection() (*mongo.Collection, error) {
	URL := os.Getenv("NAGA_DB_LOCAL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URL))

	if err != nil {
		return nil, err
	}

	collection := client.Database("grrmarchive").Collection("blogposts")

	return collection, nil
}
