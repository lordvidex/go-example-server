.PHONY: protos
server:
	go run ./cmd/server/
test:
	go test ./...
protos:
	protoc -I protos/ --go_out=. --go-grpc_out=. protos/product.proto