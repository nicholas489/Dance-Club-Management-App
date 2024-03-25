package user

import (
	"CPS406-Assignment-Backend/internal/util"
	"CPS406-Assignment-Backend/pkg/event"
	"CPS406-Assignment-Backend/pkg/jwtM"
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
		util.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user exists in the database
	var user user.User
	result := db.First(&user, "email = ?", l.Email)

	// If the user does not exist, send a JSON error message
	if result.Error != nil {
		util.SendJSONError(w, "User not found", http.StatusNotFound)
		return
	}

	// If the user exists, check if the passwords match
	if user.Password != l.Password {
		util.SendJSONError(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	privileges := util.SetPrivileges(jwtM.CustomClaims{Privileges: jwtM.Privileges{User: true}})
	tokenString, err := util.GenerateJWT(user.Email, privileges)
	if err != nil {
		util.SendJSONError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// If the passwords match, send the user details
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+tokenString)
	response := map[string]string{
		"message": "Login successful",
		"Email":   user.Email,
		"Name":    user.Name,
	}
	json.NewEncoder(w).Encode(response)
}

func PostSignup(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Parse the request body
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		util.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user already exists
	var user user.User
	result := db.First(&user, "email = ? ", u.Email)

	// If the user already exists, send an error message
	if result.Error == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	// If the user does not exist, create a new user
	db.Create(&u)
	privileges := util.SetPrivileges(jwtM.CustomClaims{Privileges: jwtM.Privileges{User: true}})
	tokenString, err := util.GenerateJWT(u.Email, privileges)
	if err != nil {
		util.SendJSONError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	// Send the user details
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+tokenString)
	json.NewEncoder(w).Encode(u)
}

func JoinEvent(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Parse the request body user + event
	var ue event.UserEvent
	err := json.NewDecoder(r.Body).Decode(&ue)
	if err != nil {
		util.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Fine the event and add the user to the event
	var e event.Event
	result := db.First(&e, "name = ?", ue.Event.Name)
	if result.Error != nil {
		util.SendJSONError(w, "Event not found", http.StatusNotFound)
		return
	}
	//update the user balance
	var u user.User
	result = db.First(&u, "email = ?", ue.User.Email)
	if result.Error != nil {
		util.SendJSONError(w, "User not found", http.StatusNotFound)
		return

	}
	u.Balance = u.Balance - e.Cost
	db.Save(&u)
	// Add the user to the event
	e.Users = append(e.Users, ue.User)
	db.Save(&e)

}

//todo: implement
