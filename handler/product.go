package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"simple_restapi/dto"
	"simple_restapi/entity"
	"strconv"
)

var ProductItems = []entity.Product{
	{
		ID:    1,
		Name:  "Nike",
		Qty:   100,
		Price: 150000,
	},
	{
		ID:    2, //  Changed ID from 1 to 2
		Name:  "Joger",
		Qty:   120,
		Price: 120000,
	},
	{
		ID:    3, //  Changed ID from 1 to 3
		Name:  "Jeans",
		Qty:   150,
		Price: 100000,
	},
}

func getNextID() int {
	maxID := 0
	for _, item := range ProductItems {
		if item.ID > maxID {
			maxID = item.ID
		}
	}
	return maxID + 1
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	resData := dto.ResponseModels{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "Success",
		Data:            ProductItems,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	json.NewEncoder(w).Encode(resData)

}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var ProductItem entity.Product

	// Decode the incoming JSON request
	err := json.NewDecoder(r.Body).Decode(&ProductItem)
	if err != nil {
		// Log the exact error to understand the problem
		log.Printf("Error decoding JSON : %v", err)
		http.Error(w, "invalid request", http.StatusBadRequest) //400
		return
	}

	// Assign a unique ID to the new product
	ProductItem.ID = getNextID()

	// Append the new product to the list of products
	ProductItems = append(ProductItems, ProductItem)

	// Prepare the response data
	resData := dto.ResponseModels{
		ResponseCode:    http.StatusCreated, // Use 201 Created for a POST request
		ResponseMessage: "Product added successfully",
		Data:            ProductItems, // Include the updated list of products
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

	for i, product := range ProductItems {
		if product.ID == updatedProduct.ID {
			ProductItems[i] = updatedProduct

			resData := dto.ResponseModels{
				ResponseCode:    http.StatusOK,
				ResponseMessage: "Product updated successfully",
				Data:            ProductItems,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(resData)
			return
		}
	}

	http.Error(w, "product not found", http.StatusNotFound)
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

	// Find and remove the product with the given ID
	for index, item := range ProductItems {
		if item.ID == productID {
			// Remove the product by slicing the array
			ProductItems = append(ProductItems[:index], ProductItems[index+1:]...)

			// Prepare success response
			resData := dto.ResponseModels{
				ResponseCode:    http.StatusOK,
				ResponseMessage: "Product deleted successfully",
				Data:            ProductItems,
			}

			w.Header().Set("Context-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resData)
			return
		}

		// If the product with the given ID is not found, return an error
		http.Error(w, "product not found", http.StatusNotFound)
	}
}
