# go-example-server <a href="https://github.com/lordvidex/go-example-server"><img src="https://pkg.go.dev/badge/github.com/lordvidex/go-example-server.svg" alt="Go Reference"></a>  
Learning golang for servers and microservices

## How To Run?
The server and client processes can be started with the following make commands.  

**REST and gRPC servers**
```bash
make server # starts the REST & gRPC server
```
**gRPC client** *(get single item)*
```bash
bash ./client.sh # returns a single item through gRPC
```
**gRPC client** *(get all item)*
```bash
bash ./client.sh -all # returns a stream of items through gRPC
```
