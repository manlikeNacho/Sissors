package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manlikeNacho/Sissors/src/controller"
	"github.com/manlikeNacho/Sissors/src/repository/sliceRepo"
	"log"
)

func main() {
	//Inittailize db
	client := sliceRepo.New()
	ctrl := controller.New(client)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my url shortener",
		})
	})
	r.POST("/Url", ctrl.CreateShortUrl)
	if err := r.Run(":8080"); err != nil {
		log.Printf("Server crashed due to, %v", err)
	}
}
