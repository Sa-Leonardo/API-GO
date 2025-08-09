package usecase

import (
	"API-go/model"
	"API-go/repository"
)

type ProductUsecase struct {
	//repository
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}


func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreatProduct(product model.Product) (model.Product, error) {
	productId , err := pu.repository.CreatProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil 
}


func (pu *ProductUsecase) GetProductsById(id_product int) (*model.Product, error) {

	//chamando o repository
	product, err := pu.repository.GetProductsById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}