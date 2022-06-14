package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/lordvidex/go-example-server/data"
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
	log.Printf("I received a request of %s", r.GetId())
	return &protos.ProductResponse{
		Id:          "12",
		Name:        "This is who I am",
		Description: "This is what I'm saying",
	}, nil
}

func setupHTTPRouter() http.Handler {
	router := http.NewServeMux()
	router.Handle("/product", handlers.NewProduct(data.Product{}))
	router.Handle("/order", handlers.NewOrder())
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
