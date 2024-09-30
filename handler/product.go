package handler

import (
	"encoding/json"
	"net/http"
	"simple_restapi/dto"
	"simple_restapi/entity"
)

var ProductItems = []entity.Product{
	{
		ID:    1,
		Name:  "Nike",
		Qty:   100,
		Price: 150000,
	},
	{
		ID:    1,
		Name:  "Joger",
		Qty:   120,
		Price: 120000,
	},
	{
		ID:    1,
		Name:  "Jeans",
		Qty:   150,
		Price: 100000,
	},
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	resData := dto.ResponseModels{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "Success",
		Data:            ProductItems,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	json.NewEncoder(w).Encode(resData)

}
func AddProduct(w http.ResponseWriter, r *http.Request) {
	var ProductItem entity.Product
	err := json.NewDecoder(r.Body).Decode(&ProductItem)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest) //400
	}

	ProductItem.ID = len(ProductItems) + 1
	ProductItems = append(ProductItems, ProductItem)
	resData := dto.ResponseModels{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "Success",
		Data:            ProductItems,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resData)

}
