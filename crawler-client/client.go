package crawler

import (
	. "../api/crawler"
	"gopkg.in/resty.v0"
)

type Client struct {
	baseUrl string
}

func NewCliente(host string) (ICrawlerApi) {
	return &Client{
		baseUrl: host + "/" + ContestPath,
	}
}

func (client *Client) Start() (*Status, error) {
	var status Status
	var err error

	if _, err = resty.R().
		SetResult(&status).
		Get(client.baseUrl + "/start"); err != nil {
		return nil, err
	}

	return &status, nil
}

func (client *Client) Status() (*Status, error) {
	var status Status
	var err error

	if _, err = resty.R().
		SetResult(&status).
		Get(client.baseUrl + "/status"); err != nil {
		return nil, err
	}

	return &status, nil
}

func (client *Client) Stop() (*Status, error) {
	var status Status
	var err error

	if _, err = resty.R().
		SetResult(&status).
		Get(client.baseUrl + "/stop"); err != nil {
		return nil, err
	}

	return &status, nil
}
