package search

import (
	"github.com/gin-gonic/gin"
	"../search-engine"
	"../search-index"
	"../api"
)

type Config struct {
	indexUrl string
}

func NewConfig(indexUrl string) (*Config) {
	return &Config{
		indexUrl : indexUrl,
	}
}

type Server struct {
	searchResource *Resource
	address        string
}

func NewServer(config *Config) (*Server, error) {
	searchEngine, err := search.NewSearchEngine(search.ELASTIC, config.indexUrl)
	if err != nil {
		return nil, err
	}
	index := index.NewIndexService(searchEngine)
	service := NewService(index)
	resource := NewResource(service)

	return newServer(resource), nil
}

func newServer(searchResource *Resource) *Server {
	return &Server{
		searchResource: searchResource,
		address: ":" + api.Port,
	}
}

func (server *Server) Migrate() (error) {
	return nil
}

func (server *Server) Run() (error) {
	router := gin.Default()

	baseGroup := router.Group(api.SearchContestPath)
	{
		baseGroup.GET("/search/:name", server.searchResource.Search)
	}

	router.Run(server.address)
	return nil
}
