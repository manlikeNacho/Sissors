package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manlikeNacho/Sissors/src/controller"
)

func AuthRoutes(router *gin.Engine, ctrl *controller.Controller) {
	// panic("Implement me")
	router.GET("/user/signup", ctrl.Signup)
	router.POST("user/login", ctrl.Login)
}
