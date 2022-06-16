.PHONY: pb
server:
	go run ./cmd/server/
test:
	go test ./...
pb:
	protoc -I internal/pb/ --go_out=. --go-grpc_out=. internal/pb/product.proto