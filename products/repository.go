package products

import (
	"encoding/json"
	"io/ioutil"
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

// NewRepository returns a repository instance
func NewRepository() *repository {
	return &repository{}
}

func init() {
	// load JSON data from data/mock.json
	err := loadData()
	if err != nil {
		panic("An error occured loading json data: " + err.Error())
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
