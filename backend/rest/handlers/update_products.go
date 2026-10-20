package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/altr-r/implementing-go-backend/database"
	"github.com/altr-r/implementing-go-backend/utils"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pid, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Not a valid JSON")
		http.Error(w, "Give valid json", 400)
		return
	}

	newProduct.Id = pid

	database.UpdateProduct(newProduct)

	utils.SendData(w, "Succesfully updated product", 201)
}
