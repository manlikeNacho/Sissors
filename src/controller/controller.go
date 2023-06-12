package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/manlikeNacho/Sissors/src/repository"
)

type controller struct {
	repo repository.Repository
}

func New(repo repository.Repository) controller {
	return controller{
		repo: repo,
	}
}

func (ct controller) CreateShortUrl(c *gin.Context) {

}
