package handlers

import (
	"github.com/lordvidex/go-example-server/data"
	"net/http"
)

type Product struct {
	Data data.Product
}

func NewProduct(product data.Product) *Product {
	return &Product{
		Data: product,
	}
}

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := p.Data.ToJSON(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
}
