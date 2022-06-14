.PHONY: protos
server:
	go run ./cmd/server/

protos:
	protoc -I protos/ --go_out=. --go-grpc_out=. protos/product.proto