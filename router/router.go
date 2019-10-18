package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r :=gin.New()
	api := r.Group("/api/v1")
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"data":"pong",
		})
	})
	return r
}
