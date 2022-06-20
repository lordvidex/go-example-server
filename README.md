# go-example-server <a href="https://github.com/lordvidex/go-example-server"><img src="https://pkg.go.dev/badge/github.com/lordvidex/go-example-server.svg" alt="Go Reference"></a>  
Learning golang for servers and microservices

## How To Run?  
### Server
The server processes can be started `locally` with the make command below.  

**REST and gRPC servers**
```bash
make server # starts the REST & gRPC server
```
**Docker**
```bash
docker-compose up
```


### Client
**HTTP client**
```bash
curl http://localhost:3000/product # all items
curl https://localhost:3000/product/1 # single item
```
**gRPC client** *(get single item)*
```bash
bash ./client.sh # returns a single item through gRPC
```
**gRPC client** *(get all item)*
```bash
bash ./client.sh -all # returns a stream of items through gRPC
```
**or with grpcurl**
```bash
grpcurl -d '{"id": "2"}' --plaintext localhost:3000 Product.GetProduct # get single  

grpcurl --plaintext localhost:3000 Product.GetProducts # get all
```

### **DockerHub** *pull project directly from docker*
```bash
docker pull lordvidex/goserver:latest
```