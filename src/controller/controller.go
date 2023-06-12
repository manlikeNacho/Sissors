package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/manlikeNacho/Sissors/src/repository"
)

type Controller struct {
	repo repository.Repository
}

func New(repo repository.Repository) Controller {
	return Controller{
		repo: repo,
	}
}

func (ct Controller) CreateShortUrl(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (ct Controller) GetUrl(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
