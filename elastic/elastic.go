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

func init() {

	esCfg := elasticsearch.Config{
		Addresses: []string{
			"http://elastic:9200/",
		},
	}

	client, err := elasticsearch.NewClient(esCfg)

	utils.HandleFatalError("Error creating client", err)

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

	utils.HandleFatalError("Coundn't connect to elastic instance", err)

	log.Printf("%v", esInfo)

	// Default Mapping the string.
	err = createIndex(models.HtmlDocumentMapping)

	utils.HandleFatalError("Index creation error:", err)

}

func GetEsClient() *elasticsearch.Client {
	return client
}

func createIndex(mapping string) error {
	res, err := client.Indices.Create("index01", client.Indices.Create.WithBody(strings.NewReader(mapping)))

	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("error: %s", res)
	}
	return nil
}
