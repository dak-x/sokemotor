package main

import (
	"log"
	"net/http"
	"sokemotor/elastic"

	"github.com/gin-gonic/gin"
)

type htmlDocument struct {
	Url string `json:"url"`
	Dom string `json:"dom"`
}

func main() {

	_ = elastic.GetEsClint()

	app := gin.Default()
	app.GET("/search/:query", searchHandler)
	app.POST("/document", documentHandler)
	app.POST("/indexHtml", processHtml)
	app.Run("0.0.0.0:8080")
}

func searchHandler(c *gin.Context) {
}

func documentHandler(c *gin.Context) {

}

func processHtml(c *gin.Context) {

	var v htmlDocument

	err := c.Bind(&v)

	if err != nil {
		log.Printf("Error %v", err)
	} else {
		log.Printf("Recevied: %v", v.Url)
	}

	c.JSON(http.StatusCreated, gin.H{
		"url":     v.Url,
		"Created": "true",
	})
}
