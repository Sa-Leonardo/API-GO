package main

import (
	"API-go/controller"
	"API-go/db"
	"API-go/repository"
	"API-go/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil{
		panic(err)
	}

// CAMADA de REPOSITORY
	ProductRepository := repository.NewProductRepository(dbConnection)
	// CAMADA de USECASE
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)

	// CAMADA DE CONTROLLERS
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Está funcionando!",
		})
	})
// criando uma rota de GET para os produtos
	server.GET("/products", ProductController.GetProducts)

	// criando uma rota de POST para um produto 
	server.POST("/product", ProductController.CreatProduct)

	// criando uma rota de POST para um produto
	server.GET("/product/:productId", ProductController.GetProductsById)

	//Porta
	server.Run(":8000")
}
