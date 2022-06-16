FROM golang:1.18
WORKDIR /go/src/go-example-server
COPY . .
RUN go build -o bin/server cmd/server/server.go
CMD [ "./bin/server" ]