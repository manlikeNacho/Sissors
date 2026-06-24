package controller

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/manlikeNacho/Sissors/src/pkg/shortener"
	"github.com/manlikeNacho/Sissors/src/repository/sliceRepo"
	"github.com/manlikeNacho/Sissors/src/utils/rerrors"
)

type Controller struct {
}

func New() Controller {
	return Controller{}
}

func (ct Controller) Signup(c *gin.Context) {
	//parse request body
	var user models.SignupReq
	err := c.ShouldBindJSON(&user)
	if err != nil {
		err := rerrors.Format(rerrors.UnproccessibleEntityErr, err)
		errCode := err.(*rerrors.Err).Status()
		c.JSON(errCode, gin.H{
			"error": err,
		})
		return
	}

	err = user.ValidateUser()
	if err != nil {
		errCode := err.(*rerrors.Err).Status()
		c.JSON(errCode, gin.H{
			"error": err,
		})
		return
	}

	// checks if email already exist and returns a boolean
	if !sliceRepo.UserRepo.CheckUserExistsByEmail(user.Email) {
		err := rerrors.Format(rerrors.BadRequestErr, errors.New("email or phone_number already exist"))
		errCode := err.(*rerrors.Err).Status()
		c.JSON(errCode, gin.H{
			"error": err,
		})
		return
	}

	_, err = sliceRepo.UserRepo.SaveUser(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "signUp",
	})
}

func (ct Controller) Login(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Login",
	})
}

func (ct Controller) GetUserById(c *gin.Context) {
	userId := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"message": "user by Id...",
		"user_id": userId,
	})
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

	//generate and update short_url
	if userUrl.ShortUrl == "" {
		short_url, err := shortener.GenerateShortLink(userUrl)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "an error occured while encoding url",
			})
			return
		}
		userUrl.ShortUrl = short_url
	}

	//Check db for short_url
	if _, err = sliceRepo.UrlRepo.GetUrl(userUrl.ShortUrl); err == nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "url already exist",
		})
		return
	}

	//save url in db
	if err := sliceRepo.UrlRepo.SaveUrl(userUrl); err != nil {
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
	val, err := sliceRepo.UrlRepo.GetUrl(p)
	//if err is nil , that means url short key is already saved.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "short url doesn't exist",
		})
		return
	}

	c.JSON(http.StatusMovedPermanently, val)
}
