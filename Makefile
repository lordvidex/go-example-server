.PHONY: protos

server:
	go run ./cmd/server/server.go
client:
	go run ./cmd/client/client.go
protos:
	protoc -I protos/ --go_out=. --go-grpc_out=. protos/product.proto