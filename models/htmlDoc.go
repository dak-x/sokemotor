package models

type HtmlDocument struct {
	Url string `json:"url"`
	Dom string `json:"dom"`
}

// For elastic-search indexing
const HtmlDocumentMapping string = `"mappings": {
	"properties": {
		"url" : {"type" : ""},
		"htmltext" : {"type" : "text"},
		"time_of_registry" : {"type" : "date"},
	}
  }`
