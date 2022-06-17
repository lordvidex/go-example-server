package products

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lordvidex/go-example-server/internal/pb"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestHandler_GetSingleProductHTTP(t *testing.T) {
	testCases := []struct {
		id             string
		expectedStatus int
		expectedError  error
	}{
		{id: "1", expectedStatus: http.StatusOK},
		{id: "2", expectedStatus: http.StatusOK},
		{id: "13", expectedStatus: http.StatusNotFound},
	}
	for _, tt := range testCases {
		t.Run("When id is "+tt.id, func(t *testing.T) {
			// given
			h := &handler{
				repo: *mockRepository(),
			}
			c := makeGinMock()

			c.AddParam("id", tt.id)

			// when
			h.GetSingleProductHTTP(c)

			// then
			if c.Writer.Status() != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, c.Writer.Status())
			}
		})
	}
}

func TestHandler_GetProductsHTTP(t *testing.T) {
	h := &handler{
		repo: *mockRepository(),
	}
	respWriter := httptest.NewRecorder()
	h.GetProductsHTTP(respWriter, &http.Request{})
	if respWriter.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, respWriter.Code)
	}
	if respWriter.Body == nil || respWriter.Body.String() == "" {
		t.Error("Empty body is unexpected")
	}
}

func TestHandler_CreateProductsHTTP(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedStatus int
	}{
		{
			name:           "When body is empty JSON",
			json:           `{}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "When name is not empty",
			json:           `{"id": 15, "name": "Product 1", "description": "sadgdasf"}`,
			expectedStatus: http.StatusCreated,
		},
	}
	for _, tt := range testCases {
		h := &handler{
			repo: *mockRepository(),
		}
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.json))
			h.CreateProductsHTTP(c)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func Test_handler_GetProduct(t *testing.T) {
	type args struct {
		c context.Context
		r *pb.ProductRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.ProductResponse
		wantErr bool
	}{
		{"When id is 1", args{context.Background(), &pb.ProductRequest{Id: "1"}}, &pb.ProductResponse{Id: "1", Name: "Product 1", Description: "sadgdasf"}, false},
		{"When id is 13", args{context.Background(), &pb.ProductRequest{Id: "13"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				repo: *mockRepository(),
			}
			got, err := h.GetProduct(tt.args.c, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler_GetProducts(t *testing.T) {
	type args struct {
		srv pb.Product_GetProductsServer
	}
	var sentProducts []Product
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"fetching all products", args{srv: &mockGetProductsServer{
			Sent: &sentProducts,
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				repo: *mockRepository(),
			}

			if err := h.GetProducts(nil, tt.args.srv); (err != nil) != tt.wantErr {
				t.Errorf("GetProducts() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(sentProducts) == 0 {
				t.Error("No products were sent")
			}
		})
	}
}
