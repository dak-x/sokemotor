package models

var IndexName = "myindex"

type HtmlDocument struct {
	Url string `json:"url"`
	Dom string `json:"dom"`
}

// Mapping
var HtmlDocumentMapping string = `{
	"settings":{
		"analysis":{
		   "analyzer":{
			  "my_analyzer":{
				 "type":"custom",
				 "tokenizer":"uax_url_email",
				 "filter":[
					"lowercase",
					"asciifolding"
				 ]
			  }
		   }
		}
	 },

	"mappings": {
		"properties": {
			"lastaccessed" : {"type" : "date"},
			"url" : {"type" : "keyword"},
			"htmltext" : {
				"type" : "text",
				"analyzer":"my_analyzer"
			}
		}
	}
  }`
