package crawler

import (
	"../api"
	. "../api/crawler"
	"../sfgov-api"
	"log"
	"time"
	"../search-index"
)

const LIMIT uint = 2
const SLEEP = 2

type Worker struct {
	state  *Status
	movies *sfgov.MovieResource
	index  *index.IndexService
}

func NewWorker(state *Status, movieResource *sfgov.MovieResource, index *index.IndexService) (*Worker) {
	return &Worker{
		state        : state,
		movies       : movieResource,
		index        : index,
	}
}

func (worker *Worker) Crawler() {
	for {
		time.Sleep(SLEEP * time.Second)
		if (!worker.state.Active) {
			log.Printf("Crawler wate %d sec.", SLEEP)
			continue;
		}

		films, err := worker.craw()
		if err != nil {
			worker.state.Active = false
			log.Printf("Crawler stopped. Error: %s", err)
			continue
		}

		worker.state.Offset = worker.state.Offset + uint(films.Len())
	}
}

func (worker *Worker) craw() (*api.Films, error) {
	movies, err := worker.movies.Get(worker.state.Offset, LIMIT)
	if err != nil {
		return nil, err
	}

	films, err := convert(movies)
	if err != nil {
		return nil, err
	}

	err = worker.index.Add(films)
	if err != nil {
		return nil, err
	}

	log.Printf("Cpllected %d items", len(*movies))
	return &api.Films{}, nil
}

func convert(movies *[]sfgov.Movie) (*[]api.Film, error) {
	var films []api.Film
	for _, element := range *movies {
		films = append(films, *api.NewFilm(element.Title, element.Release_year))
	}
	return &films, nil
}
