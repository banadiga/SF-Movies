package crawler

import (
	. "../api/crawler"
	"github.com/gin-gonic/gin"
)

type Resource struct {
	service ICrawlerApi
}

func NewResource(service ICrawlerApi) (*Resource) {
	return &Resource{
		service : service,
	}
}

func (resource *Resource) Start(context *gin.Context) {
	status, err := resource.service.Start()
	responce(context, status, err)
}

func (resource *Resource) Status(context *gin.Context) {
	status, err := resource.service.Status()
	responce(context, status, err)
}

func (resource *Resource) Stop(context *gin.Context) {
	status, err := resource.service.Stop()
	responce(context, status, err)
}

func responce(context *gin.Context, status *Status, err error) {
	if err != nil {
		context.Error(err)
		context.JSON(500, gin.H{"error": err.Error()})
	} else {
		context.JSON(200, status)
	}
}
