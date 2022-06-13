package main

import (
	"context"
	"fmt"
	"github.com/lordvidex/go-example-server/data"
	"github.com/lordvidex/go-example-server/handlers"
	"github.com/lordvidex/go-example-server/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

const (
	HttpPort = "8080"
	GrpcPort = "50051"
)

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

	// create a wait group
	wg := sync.WaitGroup{}
	wg.Add(2) // for two goroutines

	// HTTP methods for router
	router := setupHTTPRouter()

	// GRPC methods and routes
	grpcServer := setupGRPCServer()

	go func() {
		err := http.ListenAndServe(":"+HttpPort, router)
		defer wg.Done()
		if err != nil {
			return
		}
	}()

	go func() {
		lis, err := net.Listen("tcp", ":"+GrpcPort)
		defer wg.Done()
		if err != nil {
			fmt.Println(fmt.Errorf("error occured starting in another port"))
			os.Exit(-1)
		}
		err = grpcServer.Serve(lis)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
	wg.Wait()
}
