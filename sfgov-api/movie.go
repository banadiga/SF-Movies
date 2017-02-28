package sfgov

import (
	"github.com/SebastiaanKlippert/go-soda"
	"encoding/json"
)

const (
	movieResourceUrl string = "https://data.sfgov.org/resource/wwmu-gmzc"
)

type Movie struct {
	Title              string `json:"title,omitempty" binding:"required"`
	Writer             string `json:"writer,omitempty" binding:"required"`
	Release_year       int    `json:"release_year,string,omitempty" binding:"required"`
	Production_company string `json:"production_company,omitempty" binding:"required"`
	Director           string `json:"director,omitempty" binding:"required"`
	Locations          string `json:"locations,omitempty" binding:"required"`
	Actor_1            string `json:"actor_1,omitempty" binding:"required"`
	Actor_2            string `json:"actor_2,omitempty" binding:"required"`
	Actor_3            string `json:"actor_3,omitempty" binding:"required"`
}

type MovieResource struct {
	url      string
	appToken string
}

func NewMovieResource(appToken string) *MovieResource {
	return &MovieResource{
		url: movieResourceUrl,
		appToken: appToken,
	}
}

func (movieResource *MovieResource) Get(offset uint, limit uint) (*[]Movie, error) {
	request := soda.NewGetRequest(movieResource.url, movieResource.appToken)
	request.Query.Limit = limit
	request.Query.Offset = offset

	request.Query.AddOrder("release_year", true)
	request.Query.AddOrder("title", true)

	response, err := request.Get()
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(response.Body)
	defer response.Body.Close()

	// read open bracket
	_, err = decoder.Token()
	if err != nil {
		return nil, err
	}

	movies, err := decodeMovies(decoder)
	return movies, err
}

func decodeMovies(decoder *json.Decoder) (*[]Movie, error) {
	var movies []Movie

	for decoder.More() {
		var movie Movie

		err := decoder.Decode(&movie)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}
	return &movies, nil
}
