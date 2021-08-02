package elastic

import (
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)

var client *elasticsearch.Client

func init() {

	esCfg := elasticsearch.Config{
		Addresses: []string{
			"http://elastic:9200/",
		},
	}

	client, err := elasticsearch.NewClient(esCfg)

	if err != nil {
		log.Fatalf("Cannot create Client: %v", err)
	}

	esInfo, err := client.Info()

	var numRetries int = 0

	for err != nil {
		// Wait for the instane to get up and running.
		time.Sleep(3 * time.Second)
		numRetries++
		if numRetries == 5 {
			break
		}
		esInfo, err = client.Info()
	}

	if err != nil {
		log.Fatalf("Elastic is not running %s", err)
	}

	log.Printf("ELASTIC IS UP: %v\n", esInfo)
	log.Printf("STARTING SOKEMOTOR SERVICE \n")

}

func GetEsClint() *elasticsearch.Client {
	return client
}
