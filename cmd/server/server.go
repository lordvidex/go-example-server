package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/lordvidex/go-example-server/handlers"
	"github.com/lordvidex/go-example-server/protos"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c" // needed for allowing http and grpc on the same port
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	HttpPort = "8081"
)

type app struct {
	httpRouter http.Handler
	grpcRouter http.Handler
}

func NewApp(http http.Handler, grpc http.Handler) app {
	return app{
		http, grpc,
	}
}

func (a app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
		a.grpcRouter.ServeHTTP(w, r)
	} else {
		a.httpRouter.ServeHTTP(w, r)
	}
}

type productProtoController struct {
	protos.UnimplementedProductServer
}

func (productProtoController) GetProduct(c context.Context, r *protos.ProductRequest) (*protos.ProductResponse, error) {
	return handlers.GetProductsGRPC()
}

func setupHTTPRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/product", handlers.GetProductsHTTP)
	// router.HandleFunc("/order", handlers.GetOrders)
	return router
}

func setupGRPCServer() (srv *grpc.Server) {
	srv = grpc.NewServer()
	pr := productProtoController{}
	protos.RegisterProductServer(srv, pr)
	reflection.Register(srv)
	return
}

func main() {
	// create new app
	app := NewApp(setupHTTPRouter(), setupGRPCServer())

	// HTTP server
	err := http.ListenAndServe(":"+HttpPort, h2c.NewHandler(app, &http2.Server{}))
	if err != nil {
		log.Fatal(err)
	}
}
