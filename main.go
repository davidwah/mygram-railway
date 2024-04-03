package main

import (
	"github.com/gin-gonic/gin"
	"mygram-railway/database"
	"mygram-railway/handler"
	"mygram-railway/repository"
	"mygram-railway/service"
)

func main() {
	r := gin.Default()
	postgres := database.NewPostgres()
	if postgres.Err != nil {
		panic(postgres.Err)
	}

	productRepo := repository.NewProductRepo(postgres)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	r.GET("/product/:id", productHandler.FindOneProduct)

	r.Run(":8080")

}
