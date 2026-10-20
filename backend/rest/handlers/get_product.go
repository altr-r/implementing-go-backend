package handlers

import (
	"net/http"
	"strconv"

	"github.com/altr-r/implementing-go-backend/database"
	"github.com/altr-r/implementing-go-backend/utils"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Not a valid id", 400)
		return
	}

	product := database.GetSingleProduct(id)

	if product == nil {
		utils.SendError(w, "product not found", 404)
		return
	}

	utils.SendData(w, product, 200)
}
