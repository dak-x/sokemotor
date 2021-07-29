package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

func main() {

	_, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal("Elastic Instance Not Running Exiting...")
	}

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

}
