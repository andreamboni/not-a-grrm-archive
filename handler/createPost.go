package handler

import (
	"context"
	b64 "encoding/base64"
	"net/http"

	"github.com/andreamboni/not-a-grrm-archive/config"
	"github.com/andreamboni/not-a-grrm-archive/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveBlogPostHandler(ctx *gin.Context) {
	request := CreateBlogPostRequest{}
	logger := config.NewLogger("SaveBlogPostHandler")
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
		logger.Errorf("error getting db connection %v", err)
	}

	if !postExists(blogPost.ID, coll) {
		coll.InsertOne(context.TODO(), blogPost)
		SendSuccess(ctx, "Blogpost archived", blogPost)
	} else {
		SendError(ctx, http.StatusForbidden, "Blogpost with title: '"+blogPost.Title+"' already exists")
	}
}

func createBlogHash(title string, theDate string, url string) string {
	return b64.StdEncoding.EncodeToString([]byte(title + theDate + url))
}

func postExists(blogId string, collection *mongo.Collection) bool {
	logger := config.NewLogger("postsExists")
	query := bson.M{"id": blogId}
	cursor, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Errorf("error trying to check if blog post exist: %v", err)
	}

	results := []bson.M{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		logger.Errorf("error trying to add results: %v", err)
	}

	return len(results) > 0
}
