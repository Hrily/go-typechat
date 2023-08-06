package models

import "github.com/hrily/go-typechat"

// SearchRequest to search for books
// The user can specify one or more of the following fields
type SearchRequest struct {
	// Title will find any books with the given title
	Title *string `json:"title,omitempty"`
	// Author will find any books with the given author
	Author *string `json:"author,omitempty"`
	// Subject will find any books about the given subject
	// eg: "tennis rules" will find books about "tennis" and "rules"
	Subject *string `json:"subject,omitempty"`
	// Query will find any books with the given query
	// Is used when the user input does not correspond to any of the other fields
	Query *string `json:"query,omitempty"`
}

func (s *SearchRequest) Validate() error {
	if s.Title == nil && s.Author == nil && s.Subject == nil && s.Query == nil {
		return &typechat.ValidationError{
			Message: "At least one of title, author, subject, or query must be specified",
		}
	}
	return nil
}
