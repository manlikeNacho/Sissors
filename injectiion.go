package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manlikeNacho/Sissors/src/controller"
	"github.com/manlikeNacho/Sissors/src/repository/sliceRepo"
	"github.com/manlikeNacho/Sissors/src/routes"
)

func Inject(d *sliceRepo.Db) (*gin.Engine, error) {
	router := gin.Default()

	ctrl := controller.New()
	routes.AuthRoutes(router, &ctrl)
	routes.UrlRoutes(router, &ctrl)
	routes.UserRoutes(router, &ctrl)
	return router, nil
}
