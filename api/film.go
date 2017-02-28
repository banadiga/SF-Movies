package api

type Film struct {
	Title string  `json:"title,omitempty" binding:"required"`
	Year  int `json:"year,omitempty" binding:"required"`
}

func NewFilm(title string, year int) (*Film) {
	return &Film{
		Title: title,
		Year: year,
	}
}

type Films struct {
	films *[]Film  `json:"films" binding:"required"`
}

func NewFilms(films *[]Film) (*Films) {
	return &Films{
		films: films,
	}
}

func (films *Films) Len() (int) {
	return len(*films.films)
}
