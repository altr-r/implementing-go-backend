package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/altr-r/implementing-go-backend/database"
	"github.com/altr-r/implementing-go-backend/utils"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var newLogin ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newLogin)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Not a valid json", http.StatusBadRequest)
		return
	}

	user := database.FindUser(newLogin.Email, newLogin.Password)

	if user == nil {
		http.Error(w, "failed to login - invalid credintials", http.StatusBadRequest)
		return
	}

	utils.SendData(w, user, http.StatusCreated)
}
