package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/hrily/go-typechat/examples/openlibrary/models"
	"github.com/pkg/errors"
)

const (
	openLibraryURL = "https://openlibrary.org"
	searchPath     = "/search.json"
	searchURL      = openLibraryURL + searchPath
)

type Client interface {
	Search(ctx context.Context, request *models.SearchRequest) (*models.SearchResponse, error)
}

type client struct{}

func New() Client {
	return &client{}
}

func (c *client) Search(
	ctx context.Context, request *models.SearchRequest,
) (*models.SearchResponse, error) {
	params := url.Values{}
	params.Add("sort", "rating")
	if request.Title != nil {
		params.Add("title", *request.Title)
	}
	if request.Author != nil {
		params.Add("author", *request.Author)
	}
	if request.Subject != nil {
		params.Add("subject", *request.Subject)
	}
	if request.Query != nil {
		params.Add("q", *request.Query)
	}

	u, _ := url.Parse(searchURL)
	u.RawQuery = params.Encode()

	fmt.Println("Searching: ", u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get search response")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read search response body")
	}

	response := &models.SearchResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse search response")
	}

	return response, nil
}
