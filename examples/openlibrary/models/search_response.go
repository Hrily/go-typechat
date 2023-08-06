package models

type SearchResponse struct {
	Books []*Book `json:"docs,omitempty"`
}

type Book struct {
	Title   string   `json:"title,omitempty"`
	Authors []string `json:"author_name,omitempty"`
}
