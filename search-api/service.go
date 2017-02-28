package search

import (
	"../api"
	. "../api/search"
	"../search-index"
)

type Service struct {
	index *index.IndexService
}

func NewService(index *index.IndexService) (ISearchApi) {
	return &Service{
		index : index,
	}
}

func (service *Service) Search(name string) (*api.Films, error) {
	service.index.Search(name)

	var films *[]api.Film
	return api.NewFilms(films), nil
}
