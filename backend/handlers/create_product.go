package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/altr-r/implementing-go-backend/database"
	"github.com/altr-r/implementing-go-backend/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Not a valid JSON")
		http.Error(w, "Give valid json", 400)
		return
	}

	newProduct.Id = len(database.Products) + 1

	database.Products = append(database.Products, newProduct)
	utils.SendData(w, database.Products, 201)
}
