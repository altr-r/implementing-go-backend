package handlers

import (
	"net/http"

	"github.com/altr-r/implementing-go-backend/database"
	"github.com/altr-r/implementing-go-backend/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.ListProduct(), 200)
}
