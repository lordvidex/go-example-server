package data

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	products []*Product
)

// Product struct for handling conversions to and from json
// while communicating with the REST interfaces
type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (p *Product) FromJSON(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(&products)
}

func (p *Product) ToJSON(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(products)
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
	file, err := os.Open(filepath.Join("data", "mock.json")) // open file
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
