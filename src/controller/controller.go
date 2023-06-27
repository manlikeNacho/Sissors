package controller

import (
	"log"
	"net/http"
	"net/url"

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

// CreateShortUrl godoc
// @Summary      Create short url
// @Description  generates short url and stores old url
// @Tags         Url
// @Accept       application/json
// @Produce      application/json
// @Success       200 {object}  models.Url{}
// @Router        /url [post]
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
	u, err := url.ParseRequestURI(userUrl.Url)
	if err != nil || u.Scheme != "https" || u.Host == "" || u.Path == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid URL, provide an absolute URL with a scheme (only https is allowed) and a host (e.g. https://example.com/path/to/resource))",
		})
		return
	}

	_, err = ct.repo.GetUrl(userUrl.ShortUrl)
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
		log.Printf("err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "an error occured while saving url",
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

// GetUrl godoc
// @Summary      Generate long url
// @Description  generates long url and redirects request
// @Tags         Url
// @Accept       application/json
// @Produce      application/json
// @Success       301 string  "https:\\google.com"
// @Router        /short_url/:short_url [get]
func (ct Controller) GetUrl(c *gin.Context) {
	p := c.Param("short_url")
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
