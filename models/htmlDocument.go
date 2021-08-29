package models

var HtmlIndexName = "myindex"

type HtmlDocument struct {
	Url          string `json:"url"`
	Dom          string `json:"dom"`
	LastAccessed string `json:"lastaccessed"`
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
