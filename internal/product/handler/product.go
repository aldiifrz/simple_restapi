package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"simple_restapi/internal/product/dto"
	"simple_restapi/internal/product/entity"
	"simple_restapi/internal/product/repository"
	"strconv"
)

var productRepo = repository.NewInMemoryProductRepository() // Initialize repository

func GetProduct(w http.ResponseWriter, r *http.Request) {
	products := productRepo.GetAll()

	resData := dto.ResponseModels{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "Success",
		Data:            products,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	json.NewEncoder(w).Encode(resData)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product

	// Decode the incoming JSON request
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		// Log the exact error to understand the problem
		log.Printf("Error decoding JSON : %v", err)
		http.Error(w, "invalid request", http.StatusBadRequest) //400
		return
	}

	product = productRepo.Add(product)

	// Prepare the response data
	resData := dto.ResponseModels{
		ResponseCode:    http.StatusCreated, // Use 201 Created for a POST request
		ResponseMessage: "Product added successfully",
		Data:            productRepo.GetAll(),
	}

	// Set the response headers and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Encode the response as JSON and send it back
	json.NewEncoder(w).Encode(resData)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updatedProduct entity.Product

	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	err = productRepo.Update(updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	resData := dto.ResponseModels{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "Product updated successfully",
		Data:            productRepo.GetAll(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resData)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Get the product ID from the URL query parameter
	productIDStr := r.URL.Query().Get("id")

	// Convert the product ID from string to integer
	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productID <= 0 {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}
	err = productRepo.Delete(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Prepare success response
	resData := dto.ResponseModels{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "Product deleted successfully",
		Data:            productRepo.GetAll(),
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resData)
	return
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productID <= 0 {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := productRepo.FindByID(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resData := dto.ResponseModels{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "Product found",
		Data:            product,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resData)
}
