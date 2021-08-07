package elastic

import (
	"fmt"
	"log"
	"sokemotor/models"
	"sokemotor/utils"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)

var client *elasticsearch.Client

// One Time Setup on init
func init() {

	esCfg := elasticsearch.Config{
		Addresses: []string{
			"http://elastic:9200/",
		},
	}
	var err error

	// Get a new client
	client, err = elasticsearch.NewClient(esCfg)

	utils.HandleFatalError("Error creating client", err)

	esInfo, err := client.Info()

	var numRetries int = 0
	// Retry till instance is up.
	for err != nil {
		// Wait for the instane to get up and running.
		time.Sleep(3 * time.Second)
		numRetries++
		if numRetries == 10 {
			break
		}
		esInfo, err = client.Info()
	}
	utils.HandleFatalError("Coundn't connect to elastic instance", err)
	// Connected to elaticSearch
	log.Printf("%v", esInfo)

	err = createIndex(models.HtmlDocumentMapping)
	utils.HandleFatalError("Index Creation Error", err)
}

func Client() *elasticsearch.Client {
	return client
}

func createIndex(mapping string) error {

	res, err := client.Indices.Exists([]string{models.IndexName})

	if err != nil {
		return err
	}

	// Index not found, so Creating one.
	if res.IsError() {
		res, err = client.Indices.Create("myindex", client.Indices.Create.WithBody(strings.NewReader(mapping)))
		log.Println("INDEX CREATED")
	}

	if err != nil {
		return err
	} else if res.IsError() {
		return fmt.Errorf("error: %s", res)
	}

	return nil
}
