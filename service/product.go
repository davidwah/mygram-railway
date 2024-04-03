package service

import (
	"mygram-railway/core"
	"mygram-railway/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepo
}

func NewProductService(productRepo *repository.ProductRepo) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) FindOneProduct(id int) (*core.FindOneProductResponse, error) {
	product, err := s.productRepo.FindProductByID(uint(id))
	if err != nil {
		return nil, err
	}

	resp := &core.FindOneProductResponse{
		ProductName: product.Name,
		Quantity:    product.Quantity,
	}

	return resp, nil
}
