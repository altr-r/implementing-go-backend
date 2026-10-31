package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/altr-r/implementing-go-backend/database"
	"github.com/altr-r/implementing-go-backend/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// parse jwt
	// parse header and payload or claims
	// hmac-sha-256 algorithm -> hash(header, payload, secret key)
	// parse signature part from the jwt
	// if the signature and hash is same => forward to create products
	// otherwise 401 status code with Unauthorized
	
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Not a valid JSON")
		http.Error(w, "Give valid json", 400)
		return
	}

	createdProduct := database.StoreProducts(newProduct)

	utils.SendData(w, createdProduct, 201)
}
