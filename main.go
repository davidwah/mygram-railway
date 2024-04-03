package main

import (
	"github.com/gin-gonic/gin"
	"mygram-railway/database"
	"mygram-railway/handler"
	"mygram-railway/repository"
	"mygram-railway/service"
	"os"
)

func main() {
	r := gin.Default()
	DB := database.NewPostgres()
	if DB.Err != nil {
		panic(DB.Err)
	}

	productRepo := repository.NewProductRepo(DB)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// user
	userRepo := repository.NewUserRepo(DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.GET("/product/:id", productHandler.FindOneProduct)
	r.GET("/user/:id", userHandler.FindOneUser)

	r.Run(":" + os.Getenv("PORT"))

}
