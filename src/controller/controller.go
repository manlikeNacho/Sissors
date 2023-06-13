package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/manlikeNacho/Sissors/src/pkg/shortener"
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
	//Parse request body
	var userUrl *models.Url
	if err := c.ShouldBindJSON(&userUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}
	//auth url
	// r, err := url.Parse(userUrl.Url)

	_, err := ct.repo.GetUrl(userUrl.ShortUrl)
	//if err is nil , that means url short key is already saved.
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "url already exist",
		})
		return
	}
	//generate short url
	short_url, err := shortener.GenerateShortLink(userUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "an error occured while encoding url",
		})
		return
	}

	//Update userUrl
	userUrl.ShortUrl = short_url
	//save url
	if err := ct.repo.SaveUrl(userUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("An error occured while saving url, %v", err),
			"url":   userUrl.Url,
		})
		return
	}
	//return
	c.JSON(http.StatusOK, gin.H{
		"message":     "short url created",
		"orignal_url": userUrl.Url,
		"short_url":   userUrl.ShortUrl,
	})
}

func (ct Controller) GetUrl(c *gin.Context) {
	p := c.Param("short_url")
	//auth url
	// r, err := url.Parse(userUrl.Url)

	val, err := ct.repo.GetUrl(p)
	//if err is nil , that means url short key is already saved.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "short url doesn't exist",
		})
		return
	}

	c.JSON(http.StatusMovedPermanently, val)
}
