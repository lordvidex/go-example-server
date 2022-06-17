package products

import (
	"github.com/gin-gonic/gin"
	"github.com/lordvidex/go-example-server/internal/pb"
	"google.golang.org/grpc"
	"net/http"
	"net/http/httptest"
	"strconv"
)

type mockGetProductsServer struct {
	grpc.ServerStream
	Sent *[]Product
}

func (m mockGetProductsServer) Send(response *pb.ProductResponse) error {

	var err error
	*m.Sent = append(*m.Sent, Product{
		Id:          func() (x int) { x, err = strconv.Atoi(response.Id); return }(),
		Name:        response.Name,
		Description: response.Description,
	})
	return err
}

// makeGinMock returns a mock gin context
func makeGinMock() *gin.Context {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = &http.Request{}
	return c
}

// mockRepository returns a stub of our repository struct with 3 products
func mockRepository() *repository {
	r := &repository{}
	_, _ = r.AddProduct(Product{Id: 1, Name: "Product 1", Description: "sadgdasf"})
	_, _ = r.AddProduct(Product{Id: 2, Name: "Product 2", Description: "sadgdasf"})
	_, _ = r.AddProduct(Product{Id: 3, Name: "Product 3", Description: "sadgdasf"})
	return r
}
