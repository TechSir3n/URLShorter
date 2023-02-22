package router

import (
	"github.com/gin-gonic/gin"
	"github.com/url/shorter/handlers"
)

var APIRouter = gin.Default()

func init(){
	APIRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})
	
	APIRouter.POST("/create-short-url",func(c *gin.Context){
		handler.CreateShortUrl(c)
	})

	APIRouter.GET("/:shortUrl",func(c *gin.Context){
		handler.RedirectWithShortURL(c)
	})

}