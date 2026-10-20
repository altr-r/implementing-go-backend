package handlers

import (
	"net/http"
	"strconv"

	"github.com/altr-r/implementing-go-backend/database"
	"github.com/altr-r/implementing-go-backend/utils"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pid, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	database.DeleteByID(pid)

	utils.SendData(w, "Succesfully deleted product", 201)
}
