package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lordvidex/go-example-server/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = ":3000"
var (
	all = flag.Bool("all", false, "get all products")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	fmt.Println("flag all is ", *all)

	c := pb.NewProductClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// store the result
	var r interface{}
	if *all {
		stream, err2 := c.GetProducts(ctx, &pb.Empty{})
		if err2 != nil {
			log.Fatalf("could not get products: %v", err2)
		}
		r = process(stream)
	} else {
		r, err = c.GetProduct(ctx, &pb.ProductRequest{Id: "12"})
	}
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	err = json.NewEncoder(os.Stdout).Encode(r)
	if err != nil {
		log.Fatal("failed to encode response")
	}
}

func process(stream pb.Product_GetProductsClient) (r []*pb.ProductResponse) {
	// close the stream
	defer func() {
		if err := stream.CloseSend(); err != nil {
			log.Fatal(err)
		}
	}()
	// process each item
	for {
		single, eof := stream.Recv()
		if eof != nil {
			break
		}
		r = append(r, single)
	}
	return
}
