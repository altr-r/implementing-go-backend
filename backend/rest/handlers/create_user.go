package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/altr-r/implementing-go-backend/database"
	"github.com/altr-r/implementing-go-backend/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Not a valid json", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store()

	utils.SendData(w, createdUser, http.StatusCreated)
}
