package main

import (
	"context"
	"encoding/json"
	"github.com/lordvidex/go-example-server/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := protos.NewProductClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetProduct(ctx, &protos.ProductRequest{Id: "12"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Fatal(json.NewEncoder(os.Stdout).Encode(r))

}
