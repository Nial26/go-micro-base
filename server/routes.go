package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-micro-base/item"
)

func (s *server) registerRoutes() {
	s.router.GET("/", s.helloAPI())
	s.router.POST("/item", s.createItem())
	s.router.GET("/item/:item_id", s.getItem())
}

func (s *server) helloAPI() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	}
}

func (s server) createItem() gin.HandlerFunc {
	ir := item.NewRepository(s.db)
	is := item.NewService(ir)
	return func(context *gin.Context) {
		var req item.Item
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			context.Abort()
		}
		ci, err := is.CreateItem(req)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			context.Abort()
		}
		context.JSON(http.StatusOK, gin.H{"data": ci})
	}
}


func (s server) getItem() gin.HandlerFunc {
	ir := item.NewRepository(s.db)
	is := item.NewService(ir)
	return func(context *gin.Context) {
		id := context.Param("item_id")
		i, err := is.FindItem(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			context.Abort()
		}
		context.JSON(http.StatusOK, gin.H{"data" : i})
	}
}
