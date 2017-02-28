package crawler

const (
	ContestPath string = "crawler-api"
)

type ICrawlerApi interface {
	Start() (*Status, error)
	Status() (*Status, error)
	Stop() (*Status, error)
}

type Status struct {
	Active bool `json:"active" binding:"required"`
	Offset uint `json:"offset" binding:"required"`
}

func NewStatus(active bool) (*Status) {
	return &Status{
		Active: active,
	}
}
