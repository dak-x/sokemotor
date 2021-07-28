package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	app.GET("/", searchHandler)
	app.GET("/search/:query", searchHandler)
	app.POST("/document", documentHandler)

	app.Run("0.0.0.0:8080")
}

func searchHandler(c *gin.Context) {
	c.String(http.StatusAccepted, "You Pinged")
}

func documentHandler(c *gin.Context) {

}
