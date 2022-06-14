package repository

import "github.com/lordvidex/go-example-server/data"

var (
	products []*data.Product
)

// Custom repo functions here 
func GetProducts() ([]*data.Product, error) {
	return products, nil
}