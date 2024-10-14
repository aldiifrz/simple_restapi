package main

import (
	"fmt"
	"net/http"
	"simple_restapi/handler"
	"simple_restapi/repository"
	"simple_restapi/service"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handler.GetProduct(w, r)
	case "POST":
		handler.AddProduct(w, r)
	case "PUT":
		handler.UpdateProduct(w, r)
	case "DELETE":
		handler.DeleteProduct(w, r)
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
	http.HandleFunc("/product", productHandler)
	http.Handle("/user", userHandler) // Route for user-related requests

	// Start the server
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
