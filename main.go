package main

import (
	"github.com/lordvidex/go-example-server/data"
	"github.com/lordvidex/go-example-server/handlers"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/product", handlers.NewProduct(data.Product{}))
	router.Handle("/order", handlers.NewOrder())
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
