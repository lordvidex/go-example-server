package products

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	myerrors "github.com/lordvidex/go-example-server/internal/common/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lordvidex/go-example-server/internal/pb"
	"google.golang.org/grpc"
)

// handler receives a repository
// in large applications where handlers need more than one repo
// we can use a service layer in-between handler and repository
type handler struct {
	repo repository
	pb.UnimplementedProductServer
}

func NewHandler(repo repository, grpc *grpc.Server, router *mux.Router) *handler {
	h := &handler{repo: repo}
	h.registerHTTPHandlers(router)
	pb.RegisterProductServer(grpc, h)
	return h
}

func (h *handler) registerHTTPHandlers(router *mux.Router) {
	router.HandleFunc("/{id}", h.GetSingleProductHTTP).Methods("GET")
	router.HandleFunc("", h.GetProductsHTTP).Methods("GET")
	router.HandleFunc("", h.CreateProductsHTTP).Methods("POST")
}

// GetSingleProductHTTP returns a single product through HTTP GET request
func (h *handler) GetSingleProductHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		r := myerrors.BadRequest{}
		w.WriteHeader(http.StatusBadRequest)
		_ = r.ToJSON(w)
		return
	}
	res, err := h.repo.GetProductWithId(id)
	if err != nil {
		if er, ok := err.(myerrors.NotFound); ok {
			w.WriteHeader(er.StatusCode())
			_ = er.ToJSON(w)
		} else {
			log.Fatal("Failed to parse error to JSON", err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = res.ToJSON(w)
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

func (h *handler) CreateProductsHTTP(w http.ResponseWriter, r *http.Request) {
	var product Product

	err := product.FromJSON(r.Body)
	if err != nil || !product.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		_ = (&myerrors.BadRequest{}).ToJSON(w)
		return
	}
	product, err = h.repo.AddProduct(product)
	if err != nil {
		_ = (&myerrors.InternalServerError{}).ToJSON(w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = product.ToJSON(w)
}

func (h *handler) GetProduct(_ context.Context, r *pb.ProductRequest) (*pb.ProductResponse, error) {
	i, err := strconv.Atoi(r.Id)
	if err != nil {
		return nil, err
	}
	product, err := h.repo.GetProductWithId(i)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Product not found")
	}
	return productToProto(*product), nil
}

func (h *handler) GetProducts(_ *pb.Empty, srv pb.Product_GetProductsServer) error {
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

// productToProto converts a product to a pb.Product for gRPC
func productToProto(prod Product) *pb.ProductResponse {
	return &pb.ProductResponse{
		Id:          strconv.Itoa(prod.Id),
		Name:        prod.Name,
		Description: prod.Description,
	}
}
