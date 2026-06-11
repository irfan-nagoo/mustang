package main

import (
	"fmt"
	"os"

	"github.com/mustang/pkg/core/service"
	"github.com/mustang/pkg/web"
	"github.com/mustang/pkg/db"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	docs "github.com/mustang/docs"
   	swaggerfiles "github.com/swaggo/files"
  	ginSwagger "github.com/swaggo/gin-swagger"
)


func main(){

	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Error occured while loading env file: %s", err))
	}

	dbClient := db.GetDatabaseClient()
	cs := service.CartService{
		Colln: dbClient.Database("mustang").Collection("cart"),
	}
	cr := web.CartRoute{
		Service: &cs,
	}
	router := gin.Default()
	router.POST("/carts", cr.AddToCart)
	router.GET("/carts/users/:id", cr.GetCart)
	router.DELETE("/carts/users/:id/items/:itemId", cr.DeleteFromCart)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))
}