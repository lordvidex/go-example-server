package products

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"

	"github.com/lordvidex/go-example-server/protos"
	"google.golang.org/grpc"
)

// handler receives a repository
// in large applications where handlers need more than one repo
// we can use a service layer in-between handler and repository
type handler struct {
	repo repository
	protos.UnimplementedProductServer
}

func NewHandler(repo repository, grpc *grpc.Server) *handler {
	h := &handler{repo: repo}
	protos.RegisterProductServer(grpc, h)
	return h
}

// GetProductsHTTP returns the first product we have through HTTP GET request
func (h *handler) GetProductsHTTP(w http.ResponseWriter, _ *http.Request) {
	products, err := h.repo.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
	ans := func(t []*Product) []Product {
		arr := make([]Product, len(t))
		for i, p := range t {
			arr[i] = *p
		}
		return arr
	}(products)
	err = json.NewEncoder(w).Encode(ans)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("An error occured"))
		return
	}
}

func (h *handler) GetProduct(c context.Context, r *protos.ProductRequest) (*protos.ProductResponse, error) {
	products, err := h.repo.GetProducts()
	if err != nil {
		return nil, err
	}
	return productToProto(*products[0]), nil
}

func (h *handler) GetProducts(_ *protos.Empty, srv protos.Product_GetProductsServer) error {
	products, err := h.repo.GetProducts()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	for _, p := range products {
		err = srv.Send(productToProto(*p))
		if err != nil {
			return status.Error(codes.DataLoss, err.Error())
		}
	}
	return nil
}

// productToProto converts a product to a protos.Product for gRPC
func productToProto(prod Product) *protos.ProductResponse {
	return &protos.ProductResponse{
		Id:          strconv.Itoa(prod.Id),
		Name:        prod.Name,
		Description: prod.Description,
	}
}
