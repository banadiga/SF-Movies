package search

import (
	"../api"
	. "../api/search"
	"github.com/gin-gonic/gin"
)

type Resource struct {
	service ISearchApi
}

func NewResource(service ISearchApi) (*Resource) {
	return &Resource{
		service : service,
	}
}

func (resource *Resource) Search(context *gin.Context) {
	name := context.Param("name")

	films, err := resource.service.Search(name)
	responce(context, films, err)
}

func responce(context *gin.Context, films *api.Films, err error) {
	if err != nil {
		context.Error(err)
		context.JSON(500, gin.H{"error": err.Error()})
	} else {
		context.JSON(200, films)
	}
}
