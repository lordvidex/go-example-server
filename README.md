# go-example-server <a href="https://github.com/lordvidex/go-example-server"><img src="https://pkg.go.dev/badge/github.com/lordvidex/go-example-server.svg" alt="Go Reference"></a>  
Learning golang for servers and microservices

## How To Run?

> ### Server

The server processes can be started `locally` with the make command below.  

**REST and gRPC servers**
```bash
make server # starts the REST & gRPC server
```
OR
**Docker**
```bash
docker-compose up
```

> ### Client
#### HTTP
**HTTP client**
```bash
curl http://localhost:3000/product # all items
curl https://localhost:3000/product/1 # single item
```
#### GRPC
**gRPC client** 
```bash
bash ./client.sh # returns a single item through gRPC
bash ./client.sh -all # returns a stream of items through gRPC
```
OR.   
**with grpcurl**
```bash
grpcurl -d '{"id": "2"}' --plaintext localhost:3000 Product.GetProduct # get single  

grpcurl --plaintext localhost:3000 Product.GetProducts # get all
```

### **DockerHub** *pull project directly from docker*
Project can be pulled from docker hub and run directly from docker

```bash
docker pull lordvidex/goserver:latest
```
