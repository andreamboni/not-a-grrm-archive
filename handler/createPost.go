package handler

import (
	"context"
	b64 "encoding/base64"
	"log"
	"net/http"

	"github.com/andreamboni/not-a-grrm-archive/config"
	"github.com/andreamboni/not-a-grrm-archive/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveBlogPostHandler(ctx *gin.Context) {
	request := CreateBlogPostRequest{}
	ctx.BindJSON(&request)

	blogPost := model.BlogPost{
		ID:        createBlogHash(request.Title, request.TheDate, request.URL),
		PostImage: request.PostImage,
		Title:     request.Title,
		URL:       request.URL,
		TheDate:   request.TheDate,
		Content:   request.Content,
	}

	coll, err := config.GetMongoDBCollection()

	if err != nil {
		panic(err)
	}

	if !recordExists(blogPost.ID, coll) {
		coll.InsertOne(context.TODO(), blogPost)
		ctx.Header("Content-type", "application/json")
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Blogpost archived",
			"blogPost": blogPost,
		})
	} else {
		ctx.Header("Content-type", "application/json")
		ctx.JSON(http.StatusForbidden, gin.H{
			"message":  "Blogpost already exists",
			"blogPost": blogPost,
		})
	}

}

func createBlogHash(title string, theDate string, url string) string {
	return b64.StdEncoding.EncodeToString([]byte(title + theDate + url))
}

func recordExists(blogId string, collection *mongo.Collection) bool {
	query := bson.M{"id": blogId}
	cursor, err := collection.Find(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}

	results := []bson.M{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	return len(results) > 0
}
