package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/manlikeNacho/Sissors/docs"
	"github.com/manlikeNacho/Sissors/src/controller"
	"github.com/manlikeNacho/Sissors/src/repository/sliceRepo"
	"github.com/manlikeNacho/Sissors/src/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

func inject(d *sliceRepo.Db) (*gin.Engine, error) {
	//Initailize db collections
	sliceRepo.TokenRepo.InitializeTokenDB(d.Db, "snipbit", "token")
	sliceRepo.UrlRepo.InitializeUrlDB(d.Db, "snipbit", "url")
	router := gin.Default()

	ctrl := controller.New()
	routes.AuthRoutes(router, &ctrl)
	routes.UrlRoutes(router, &ctrl)
	routes.UserRoutes(router, &ctrl)
	return router, nil
}

func main() {
	// Initialize db
	client := sliceRepo.New()

	// Inject db instance into controllers and routers
	r, err := inject(client)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r.Use(gin.Logger())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := r.Run(":8080"); err != nil {
		log.Printf("Server crashed due to, %v", err)
	}
}
