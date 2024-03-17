package user

import (
	"CPS406-Assignment-Backend/pkg/login"
	"CPS406-Assignment-Backend/pkg/user"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "text/plain")
	var users []user.User
	// get all the users from the database
	db.Find(&users)
	// send the users as a response list
	json.NewEncoder(w).Encode(users)

}

func GetUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "text/plain")
	// get the id from the url
	id := chi.URLParam(r, "id")

	// get the user from the database and save it in the user variable
	var user user.User
	db.Debug().Find(&user, "id = ?", id)

	// send the user as a response
	json.NewEncoder(w).Encode(user)
}

func PostLogin(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Parse the request body
	var l login.Login
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user exists in the database
	var user user.User
	result := db.First(&user, "username = ?", l.Email)

	// If the user does not exist, send an error message
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// If the user exists, check if the passwords match
	if user.Password != l.Password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// If the passwords match, send the user details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
