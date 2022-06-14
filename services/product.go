package services

import (
	"github.com/lordvidex/go-example-server/data"
	"github.com/lordvidex/go-example-server/repository"
)

func GetProducts() ([]*data.Product, error) {
	return repository.GetProducts()
}

func NewProduct() data.Product {
	return data.Product{}
}