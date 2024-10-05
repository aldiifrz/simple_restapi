package main

import (
	"fmt"
	"net/http"
	"simple_restapi/handler"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handler.GetProduct(w, r)
	case "POST":
		handler.AddProduct(w, r)
	case "DELETE":
		handler.DeleteProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/product", productHandler)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
