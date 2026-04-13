package handler

import (
	"fmt"
	"net/http"
)

type Order struct {}

// create method
func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Order created")
}

// list method
func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of orders")
}

// get method
func (o *Order) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get order details")
}

// update method
func (o *Order) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Order updated")
}

// delete method
func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Order deleted")
}