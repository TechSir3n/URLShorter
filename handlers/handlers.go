package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/url/shorter/database/redis"
	"github.com/url/shorter/shortener"
)


type ResponseURL struct {
	LongUrl string `json:"long_url"`
	UserID  string `json:"user_id"`
}

func CreateShortUrl(c *gin.Context) {
	var longUrl ResponseURL

	if err := c.ShouldBindJSON(&longUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(longUrl.LongUrl,longUrl.UserID)
	database.SaveUrlShort(shortUrl, longUrl.LongUrl, longUrl.UserID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Your url shorter is ready",
		"short_url": "http://localhost:8080/" + shortUrl,
	})

}

func RedirectWithShortURL(c *gin.Context) {
	url := c.Param("shortUrl")
	shortUrl := database.GetUrlShort(url)
	c.Redirect(301, shortUrl)
}
