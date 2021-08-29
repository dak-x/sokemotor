package controller

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

	utils.HandleFatalError("error creating elastic client", err)

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
	utils.HandleFatalError("cound not connect to elastic instance", err)

	// Connected to elaticSearch
	log.Printf("%v", esInfo)

	err = createIndex(models.HtmlDocumentMapping)
	utils.HandleFatalError("index Creation Error", err)
}

func Client() *elasticsearch.Client {
	return client
}

// Inserts the given htmlDocument into the ES Index.
func InsertIntoIndex(data models.HtmlDocument) error {

	dataBytes, err := json.Marshal(data)

	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:   models.HtmlIndexName,
		Body:    bytes.NewReader(dataBytes),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), client)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error inserting into index")
	}

	return nil
}

// Returns the search results for a query string.
func SearchIndex(queryString string) ([]interface{}, error) {

	// Build the request body.
	query := buildQuery(queryString)

	// Perform the search request.
	res, err := client.Search(
		client.Search.WithIndex(models.HtmlIndexName),
		client.Search.WithBody(strings.NewReader(query)),
	)

	if err != nil {
		log.Fatalf("error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var r interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("error parsing the response body: %s", err)
	}

	searchResponce := r.(map[string]interface{})

	queryResults := searchResponce["hits"].(map[string]interface{})["hits"].([]interface{})

	log.Printf("query response: %v", r)

	return queryResults, nil
}

// Creates an index with the given string.
func createIndex(mapping string) error {

	exists, err := client.Indices.Exists([]string{models.HtmlIndexName})

	if err != nil {
		return err
	}

	var res *esapi.Response = nil
	// Index not found, so Creating one.
	if exists.IsError() {
		res, err = client.Indices.Create("myindex", client.Indices.Create.WithBody(strings.NewReader(mapping)))
		log.Println("index created")
	} else {
		log.Println("index already present")
	}

	// Index Creation Error.
	if err != nil {
		return err

	} else if res != nil && res.IsError() {
		return fmt.Errorf("error: %v", res)
	}

	return nil
}

func buildQuery(queryString string) string {

	query := fmt.Sprintf(queryTemplate, queryString)
	return query
}

const queryTemplate = `{
		"query": {
		"match": { "dom": "%v" }
		},
		"highlight": {
		"fields": {
			"dom": {}
		}
		},
		"size" : 15
	}
	
   `
