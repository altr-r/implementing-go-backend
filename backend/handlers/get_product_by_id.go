package handlers

import (
	"net/http"
	"strconv"

	"github.com/altr-r/implementing-go-backend/database"
	"github.com/altr-r/implementing-go-backend/utils"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Not a valid id", 400)
		return
	}

	for _, product := range database.Products {
		if product.Id == id {
			utils.SendData(w, product, 200)
			return
		}
	}

	utils.SendData(w, "data pai nai", 404)
}
