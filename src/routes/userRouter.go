package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manlikeNacho/Sissors/src/controller"
	"github.com/manlikeNacho/Sissors/src/middlewares"
)

func UserRoutes(router *gin.Engine, ctrl *controller.Controller) {
	// router.Use(middlewares.AuthMiddleware)
	// router.GET("/user", ctrl.GetUser)
	router.GET("/user/:user_id", middlewares.AuthUser(), ctrl.GetUserById)
}
