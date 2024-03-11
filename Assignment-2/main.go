package main

import (
	_ "github.com/roihan12/assignment-2/docs"
	"github.com/gin-gonic/gin"
	"github.com/roihan12/assignment-2/database"
	"github.com/roihan12/assignment-2/router"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Orders API
//	@version		1.0
//	@description	This is a service for managing orders
//	@termsOfService	http://swagger.io/terms
//	@contact.name	API Support
//	@contact.email	roihansori12@gmail.com
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:5000
//	@BasePath		/api

const PORT = ":5000"

func main() {
	ginRouter := gin.Default()

	database.StartDB()
	db := database.DB()
	router.SetupOrderRoute(ginRouter, db)
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	ginRouter.Run(PORT)
}
