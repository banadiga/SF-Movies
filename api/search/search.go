package search

import (
	"../"
)

const (
	ContestPath string = "search-api"
)

type ISearchApi interface {
	Search(name string) (*api.Films, error)
}