package products

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type repository struct {
}

// replace with db instance fetch
var products []*Product

// custom repo functions here

func (r *repository) GetProducts() ([]*Product, error) {
	return products, nil
}

func (r *repository) AddProduct(product Product) (Product, error) {
	products = append(products, &product)
	return product, nil
}

// NewRepository returns a repository instance
func NewRepository() *repository {
	return &repository{}
}

func init() {
	// load JSON data from data/mock.json
	err := loadData()
	if err != nil {
		log.Println("mock JSON data was not loaded with stacktrace ", err.Error())
	}
}

func loadData() (err error) {
	var byteValue []byte
	file, err := os.Open(filepath.Join("products", "mock.json")) // open file
	if err != nil {
		goto END
	}

	byteValue, err = ioutil.ReadAll(file) // read bytes from file
	if err != nil {
		goto END
	}

	err = json.Unmarshal(byteValue, &products) // read to json
	if err != nil {
		goto END
	}
END:
	return
}
