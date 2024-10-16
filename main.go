package main

import (
	"fmt"
	"net/http"

	// Alias the product and user andler imports to avoid naming conflict
	productHandler "simple_restapi/internal/product/handler"
	"simple_restapi/internal/user/handler"
	"simple_restapi/internal/user/repository"
	"simple_restapi/internal/user/service"
)

func productHandlerFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		productHandler.GetProduct(w, r)
	case "POST":
		productHandler.AddProduct(w, r)
	case "PUT":
		productHandler.UpdateProduct(w, r)
	case "DELETE":
		productHandler.DeleteProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Initialize repository, service, and handler for users
	userRepo := repository.NewInMemoryUserRepository() // Repository for users
	userService := service.NewUserService(userRepo)    // Service for users
	userHandler := handler.NewUserHandler(userService) // Handler for users

	// Register the routes for products and users
	http.HandleFunc("/product", productHandlerFunc)
	http.Handle("/user", userHandler) // Route for user-related requests

	// Start the server
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
