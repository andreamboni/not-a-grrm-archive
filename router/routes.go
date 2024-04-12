package router

import (
	"github.com/andreamboni/not-a-grrm-archive/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/not-a-grrm-archive/v1"
	v1 := router.Group(basePath)
	{
		// item
		// v1.GET("/items", item.ListItemsHandler)
		// v1.GET("/itemById", item.ListItemByIdHandler)
		// v1.GET("/deletedItems", item.ListDeletedItems)
		v1.GET("/fetchAllBlogPosts", handler.FetchAllBlogPosts)
		v1.POST("/saveBlogPost", handler.SaveBlogPostHandler)
		// v1.PUT("/updateItem", item.UpdateItemHandler)
		// v1.PUT("/recoverItem", item.RecoverItemHandler)
		// v1.DELETE("/deleteItem", item.DeleteItemHandler)
	}
}
