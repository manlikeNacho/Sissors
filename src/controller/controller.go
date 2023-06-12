package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/manlikeNacho/Sissors/src/pkg/shortener"
	"github.com/manlikeNacho/Sissors/src/repository"
	"net/http"
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
	var userUrl models.Url
	if err := c.ShouldBindJSON(&userUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//auth url
	//TODO: Configure auth file/pkg
	val, err := ct.repo.GetUrl(userUrl)
	if errors.Is(err, errors.New("url not found")) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "url already exist",
			"url":   val,
		})
		return
	}
	//generate short url
	shortUrl, err := shortener.GenerateShortLink(userUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("An error occured while encoding url, %v", err),
			"url":   val,
		})
		return
	}

	//Update userUrl
	userUrl.ShortUrl = shortUrl
	//save url
	if err := ct.repo.SaveUrl(userUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("An error occured while saving url, %v", err),
			"url":   val,
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
	//TODO implement me
	panic("implement me")
}
