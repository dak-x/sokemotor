package models

type SearchResponse struct {
	Url          string `json:"url"`
	Title        string `json:"title"`
	BoxContent   string `json:"content"`
	LastAccessed string `json:"lastaccessed"`
}
