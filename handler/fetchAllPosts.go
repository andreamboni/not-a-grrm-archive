package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/andreamboni/not-a-grrm-archive/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func FetchAllBlogPosts(ctx *gin.Context) {

	coll, err := config.GetMongoDBCollection()

	if err != nil {
		panic(err)
	}

	query := bson.M{}
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}

	results := []bson.M{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "blog posts retrivied successfully",
		"data":    results,
	})

}
