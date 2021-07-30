package main

import (
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

func getEsClint() *elasticsearch.Client {

	esCfg := elasticsearch.Config{
		Addresses: []string{
			"http://elastic:9200/",
		},
	}

	es, err := elasticsearch.NewClient(esCfg)

	if err != nil {
		log.Fatalf("Cannot create Client: %v", err)
	}

	res, err := es.Info()

	var numRetry int = 0

	for err != nil {
		time.Sleep(3 * time.Second)
		numRetry++
		if numRetry == 5 {
			break
		}
		res, err = es.Info()
	}

	if err != nil {
		log.Fatalf("Elastic is not running %s", err)
	}

	log.Printf("ELASTIC IS UP: %v", res)
	log.Printf("STARTING SOKEMOTOR SERVICE")

	return es
}

func main() {

	_ = getEsClint()

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
