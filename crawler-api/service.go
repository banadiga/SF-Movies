package crawler

import (
	. "../api/crawler"
)

type CrawlerService struct {
	state *Status
}

func NewCrawlerService(state *Status) (ICrawlerApi) {
	return &CrawlerService{
		state : state,
	}
}

func (service *CrawlerService) Start() (*Status, error) {
	service.state.Active = true
	return service.state, nil
}

func (service *CrawlerService) Status() (*Status, error) {
	return service.state, nil
}

func (service *CrawlerService) Stop() (*Status, error) {
	service.state.Active = false
	return service.state, nil
}
