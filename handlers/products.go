package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lordvidex/go-example-server/data"
	"github.com/lordvidex/go-example-server/protos"
	"github.com/lordvidex/go-example-server/repository"
)

// GetProductsHTTP returns the first product we have through HTTP GET request
func GetProductsHTTP(w http.ResponseWriter, r *http.Request) {
	products, err := repository.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
	ans := func(t []*data.Product) []data.Product {
		arr := make([]data.Product, len(t))
		for i, p := range t {
			arr[i] = *p
		}
		return arr
	}(products)
	json.NewEncoder(w).Encode(ans)
}

// GetProductsGRPC returns the first product we have through a GRPC channel
func GetProductsGRPC() (*protos.ProductResponse, error) {
	products, err := repository.GetProducts()
	if err != nil {
		return nil, err
	}
	return productToProtos(*products[0]), nil
}

func productToProtos(prod data.Product) *protos.ProductResponse {
	return &protos.ProductResponse{
		Id:        strconv.Itoa(prod.Id),
		Name:      prod.Name,
		Description: prod.Description,
	}
}
