package main

import (
	"github.com/lordvidex/go-example-server/products"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"strings"

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

func NewApp(http http.Handler, grpc http.Handler) *app {
	return &app{
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

func main() {
	// grpc server
	grpcServer := setupGRPCServer()

	// http server
	router := http.NewServeMux()

	// create new product handler
	prh := products.NewHandler(*products.NewRepository(), grpcServer)

	// setup routes
	// TODO: introduce sub routers later and let prh have single sub router for all requests
	router.HandleFunc("/product", prh.GetProductsHTTP)

	// create new app
	app := NewApp(router, grpcServer)

	// HTTP server
	err := http.ListenAndServe(":"+HttpPort, h2c.NewHandler(app, &http2.Server{}))
	if err != nil {
		log.Fatal(err)
	}
}

func setupGRPCServer() (srv *grpc.Server) {
	srv = grpc.NewServer()
	reflection.Register(srv)
	return
}
