package main

import (
	"log"
	"net/http"

	elastic "sokemotor/controller"
	"sokemotor/models"

	"github.com/gin-gonic/gin"
)

func main() {

	_ = elastic.Client()

	app := gin.Default()
	app.GET("/search/:query", searchHandler)
	app.POST("/document", documentHandler)
	app.POST("/indexHtml", processHTML)
	app.Run("0.0.0.0:8080")
}

func searchHandler(c *gin.Context) {
}

func documentHandler(c *gin.Context) {

}

func processHTML(c *gin.Context) {

	var v models.HtmlDocument

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
