package search

import (
	"../api"
	. "../api/search"
	"gopkg.in/resty.v0"
)

type Client struct {
	baseUrl string
}

func NewClient(host string) (ISearchApi) {
	return &Client{
		baseUrl: host + "/" + api.SearchContestPath,
	}
}

func (client *Client) Search(name string) (*api.Films, error) {
	var films api.Films
	var err error

	if _, err = resty.R().
		SetResult(&films).
		Get(client.baseUrl + "/search/" + name); err != nil {
		return nil, err
	}

	return &films, nil
}
