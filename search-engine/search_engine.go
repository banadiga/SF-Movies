package search

import (
	"errors"
)

type SearchEngineType string

func NewSearchEngineType(searchEngineType string) (SearchEngineType) {
	return SearchEngineType(searchEngineType)
}

type ISearchEngine interface {
	IndexExists(name string) (bool, error)
	CreateIndex(name string, mapping string) (error)
	Add(index string, body interface{}) (string, error)
	Search(index string, name string, value string, limit int) (int64, error)
}

func NewSearchEngine(searchEngineType SearchEngineType, url string) (ISearchEngine, error) {
	if (searchEngineType == ELASTIC) {
		searchEngine, err := NewElasticsearchEngine(url)
		return searchEngine, err
	}

	panic(errors.New("Type is not supportet yet."))
}
