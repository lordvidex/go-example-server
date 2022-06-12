package handlers

import (
	"fmt"
	"net/http"
)

type Order struct {
}

func NewOrder() *Order {
	return &Order{}
}

func (*Order) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Orders Handler hit")
}
