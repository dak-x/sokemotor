package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sokemotor/models"
	"sokemotor/utils"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
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

	exists, err := client.Indices.Exists([]string{models.IndexName})

	if err != nil {
		return err
	}

	var res *esapi.Response
	// Index not found, so Creating one.
	if exists.IsError() {
		res, err = client.Indices.Create("myindex", client.Indices.Create.WithBody(strings.NewReader(mapping)))
		log.Println("INDEX CREATED")
	}

	// Index Creation Error.
	if err != nil {
		return err

	} else if res.IsError() {
		return fmt.Errorf("error: %s", res)
	}

	return nil
}

// Inserts the given htmlDocument into the ES Index.
func InsertIntoIndex(data models.HtmlDocument) error {

	dataBytes, err := json.Marshal(data)

	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:   "test",
		Body:    bytes.NewReader(dataBytes),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), client)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	if res.IsError() {
		return fmt.Errorf("Error Inserting into Index")
	}

	return nil
}
