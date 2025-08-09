// DEFINIR A ROTA DE PRODUCT -  PARA TRATAR A REQUISIÇÃO

package controller

import (
	"API-go/model"
	"API-go/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type productController struct {
	// Usecase
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUsecase.GetProducts()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}


func ( p *productController) CreatProduct(ctx *gin.Context) {
	
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUsecase.CreatProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}


func (p *productController) GetProductsById(ctx *gin.Context) {

	//É necessario extrair o parametro da rota para passar como parametro aqui na função GetProductsById()
	id := ctx.Param("productId") // essa função retornar uma string 

	// VERIFICANDO SE A STRING NÃO ESTÁ VAZIA
	if(id == ""){

		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}


	// VERIFICANDO SE PODE CONVERTER A STRING EM UM INT

	productId, err := strconv.Atoi(id)
	if(err != nil) {
		response := model.Response{
			Message: "Id do produto precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductsById(productId)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

