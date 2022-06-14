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

func (p *Product) FromJSON(reader io.Reader) (err error) {
	err = json.NewDecoder(reader).Decode(p)
	return
}

func (p *Product) ToJSON(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(p)
}
