package crawler

import (
	"../api"
	. "../api/crawler"
	"../sfgov-api"
	"github.com/gin-gonic/gin"
	search "../search-engine"
	index "../search-index"
)

type Config struct {
	autoStart     bool
	sfgovAppToken string
	indexType     search.SearchEngineType
	indexUrl      string
}

func NewConfig(autoStart bool, sfgovAppToken string, indexType string, indexUrl string) (*Config) {
	return &Config{
		autoStart       : autoStart,
		sfgovAppToken   : sfgovAppToken,
		indexType       : search.NewSearchEngineType(indexType),
		indexUrl        : indexUrl,
	}
}

type Server struct {
	resource *Resource
	worker   *Worker
	index    *index.IndexService
	address  string
}

func NewServer(config *Config) (*Server, error) {
	searchEngine, err := search.NewSearchEngine(config.indexType, config.indexUrl)
	if err != nil {
		return nil, err
	}
	index := index.NewIndexService(searchEngine)
	status := NewStatus(config.autoStart)
	resource := NewResource(
		NewCrawlerService(
			status,
		),
	)
	worker := NewWorker(
		status,
		sfgov.NewMovieResource(config.sfgovAppToken),
		index,
	)

	return newServer(resource, worker, index), nil
}

func newServer(resource *Resource, worker  *Worker, index *index.IndexService) (*Server) {
	return &Server{
		resource: resource,
		worker: worker,
		index: index,
		address: ":" + api.Port,
	}
}

func (server *Server) Migrate() (error) {

	return nil
}

func (server *Server) Run() (error) {
	router := gin.Default()

	baseGroup := router.Group(ContestPath)
	{
		baseGroup.GET("/start", server.resource.Start)
		baseGroup.GET("/status", server.resource.Status)
		baseGroup.GET("/stop", server.resource.Stop)
	}
	server.index.Init();
	go server.worker.Crawler()

	router.Run(server.address)
	return nil
}
