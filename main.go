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
	app.POST("/indexHTML", processHTML)
	app.Run("0.0.0.0:8080")
}

func searchHandler(c *gin.Context) {

	queryString := c.Param("query")

	log.Printf("searching for: %v\n", queryString)

	val, _ := elastic.SearchIndex(queryString)

	c.JSON(http.StatusAccepted, val)
}

func documentHandler(c *gin.Context) {
	c.String(http.StatusNotImplemented, "custom docs not yet implemented.")
}

func processHTML(c *gin.Context) {

	var v models.HtmlDocument

	err := c.Bind(&v)

	if err != nil {
		log.Printf("indexHTML Error: %v", err)
		c.String(http.StatusBadRequest, "%v\n", err)
		return
	}

	// Index into es node.
	err = elastic.InsertIntoIndex(v)

	if err != nil {
		log.Printf("indexHTML Error: %v", err)
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}

	// Acknowledge that document is created.
	c.JSON(http.StatusCreated, gin.H{
		"url":     v.Url,
		"Created": "true",
	})

}
