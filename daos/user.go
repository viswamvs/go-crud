package daos

import (
	"encoding/json"
	"go-crud/daos/models"
	"net/http"

	"gorm.io/gorm"
)

var DB *gorm.DB

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func GetUser(w http.ResponseWriter, r *http.Request) {

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
