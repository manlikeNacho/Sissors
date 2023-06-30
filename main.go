package main

import (
	"github.com/manlikeNacho/Sissors/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/manlikeNacho/Sissors/src/controller"
	"github.com/manlikeNacho/Sissors/src/repository/sliceRepo"
)

// @title           Scissors
// @version         1.0
// @description     This is a simple Url-shortener server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Iheanacho Emmanuel
// @contact.url    github.com/manlikeNacho
// @contact.email  eiheanacho52@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	//Inittailize db
	client := sliceRepo.New()
	ctrl := controller.New(client)

	defer func() {
		if err := client.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my url shortener",
		})
	})
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/url", ctrl.CreateShortUrl)
	r.GET("short_url/:short_url", ctrl.GetUrl)
	if err := r.Run(":8080"); err != nil {
		log.Printf("Server crashed due to, %v", err)
	}
}
