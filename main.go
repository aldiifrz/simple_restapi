package main

import (
	"fmt"
	"net/http"
	"simple_restapi/handler"
)

func main() {
	http.HandleFunc("/product", handler.GetProduct)
	http.HandleFunc("/product", handler.AddProduct)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
