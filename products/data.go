package products

import (
	"encoding/json"
	"io"
)

// Product struct for handling conversions to and from json
// while communicating with the REST interfaces
type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// FromJSON parses a json string and sets the fields in struct (p *Product)
func (p *Product) FromJSON(reader io.Reader) (err error) {
	err = json.NewDecoder(reader).Decode(p)
	return
}

// ToJSON converts a Product struct and writes the json string to (writer io.Writer)
func (p *Product) ToJSON(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(p)
}
