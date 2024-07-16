package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manlikeNacho/Sissors/src/controller"
)

func UrlRoutes(router *gin.Engine, ctrl *controller.Controller) {
	router.POST("/url", ctrl.CreateShortUrl)
	router.GET("short_url/:short_url", ctrl.GetUrl)
}
