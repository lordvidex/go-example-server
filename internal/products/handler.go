package products

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lordvidex/go-example-server/internal/common/errors"
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

func NewHandler(repo repository, grpc *grpc.Server, group *gin.RouterGroup) *handler {
	h := &handler{repo: repo}
	h.registerHTTPHandlers(group)
	pb.RegisterProductServer(grpc, h)
	return h
}

func (h *handler) registerHTTPHandlers(group *gin.RouterGroup) {
	group.GET("", gin.WrapF(h.GetProductsHTTP))
	group.GET(":id", h.GetSingleProductHTTP)
	group.POST("", h.CreateProductsHTTP)
}

// GetSingleProductHTTP returns a single product through HTTP GET request
func (h *handler) GetSingleProductHTTP(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic("Invalid id")
	}
	res, err := h.repo.GetProductWithId(id)
	log.Println(res, err)
	if err != nil {
		if er, ok := err.(errors.NotFound); ok {
			c.Writer.WriteHeader(er.StatusCode())
			_ = er.ToJSON(c.Writer)
		} else {
			log.Fatal("Failed to parse error to JSON", err)
		}
		return
	}
	c.JSON(http.StatusOK, res)
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

func (h *handler) CreateProductsHTTP(c *gin.Context) {
	var product Product
	err := c.BindJSON(&product)
	if err != nil {
		err = (&errors.BadRequest{}).ToJSON(c.Writer)
		if err != nil {
			log.Print("Error: ", err)
			return
		}
	}
	product, err = h.repo.AddProduct(product)
	if err != nil {
		_ = (&errors.InternalServerError{}).ToJSON(c.Writer)
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (h *handler) GetProduct(c context.Context, r *pb.ProductRequest) (*pb.ProductResponse, error) {
	products, err := h.repo.GetProducts()
	if err != nil {
		return nil, err
	}
	return productToProto(*products[0]), nil
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
