package index

import (
	"../api"
	"../search-engine"
	"github.com/yanzay/log"
)

const (
	indexName = "i-name"
	mapping = `{
                        "mappings" : {
                            "film" : {
                                "properties" : {
                                    "title" : { "type" : "string", "index" : "not_analyzed" },
                                    "year" : { "type" : "string", "index" : "not_analyzed" },
                                }
                            }
                        }
                    }`
)

type IndexService struct {
	searchEngine search.ISearchEngine
}

func NewIndexService(searchEngine search.ISearchEngine) (*IndexService) {
	return &IndexService{
		searchEngine : searchEngine,
	}
}

func (index *IndexService) Init() (error) {
	exist, err := index.searchEngine.IndexExists(indexName)

	if err != nil {
		return err
	}

	if (exist) {
		log.Printf("Index with name %s exist.", indexName)
		return nil
	}

	log.Printf("Creating index with name %s.", indexName)
	if err := index.searchEngine.CreateIndex(indexName, mapping); err != nil {
		return err
	}

	return nil
}

func (index *IndexService)  Add(films *[]api.Film) (error) {
	for _, element := range *films {
		id, err := index.searchEngine.Add(indexName, element)
		if err != nil {
			log.Printf("Film `%s` does not added to index `%s`. Cose `%s`", element.Title, indexName, err.Error())
		} else {
			log.Printf("Film `%s` added to index with id %s", element.Title, id)
		}
	}
	return nil
}

func (index *IndexService) Search(name string) {
	count, err := index.searchEngine.Search(indexName, "name", name, 5)
	if err != nil {
		log.Printf("Error cause: %s", err.Error())
	} else {
		log.Printf("Result %d", count)
	}
}