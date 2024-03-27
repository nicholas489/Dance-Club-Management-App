package user

import (
	"CPS406-Assignment-Backend/internal/util"
	"CPS406-Assignment-Backend/pkg/event"
	"CPS406-Assignment-Backend/pkg/jwtM"
	"CPS406-Assignment-Backend/pkg/login"
	"CPS406-Assignment-Backend/pkg/user"
	"encoding/json"
	"fmt"
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
		util.SendJSONError(w, "Invalid password"+"password given: "+l.Password, http.StatusUnauthorized)
		return
	}
	privileges := util.SetPrivileges(jwtM.CustomClaims{Privileges: jwtM.Privileges{User: true}})
	tokenString, err := util.GenerateJWT(user.Email, privileges)
	util.SetTokenAsCookie(w, tokenString)
	if err != nil {
		util.SendJSONError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// If the passwords match, send the user details
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Login successful",
		"email":   user.Email,
		"name":    user.Name,
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
	util.SetTokenAsCookie(w, tokenString)

	// Send the user details and token as a response and set status code to 201 (Created)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Signup successful",
		"email":   u.Email,
		"name":    u.Name,
	}
	json.NewEncoder(w).Encode(response)
}

func JoinEvent(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var ue event.UserEvent
	if err := json.NewDecoder(r.Body).Decode(&ue); err != nil {
		util.SendJSONError(w, err.Error()+"ee", http.StatusBadRequest)
		return
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var e event.Event
	if result := tx.Preload("Users").Where("name = ?", ue.EventName).First(&e); result.Error != nil {
		tx.Rollback()
		util.SendJSONError(w, "Event not found", http.StatusNotFound)
		return
	}

	var u user.User
	if result := tx.Where("email = ?", ue.UserEmail).First(&u); result.Error != nil {
		tx.Rollback()
		util.SendJSONError(w, "User not found", http.StatusNotFound)
		return
	}

	for _, user := range e.Users {
		if user.Email == u.Email {
			tx.Rollback()
			util.SendJSONError(w, "User already in event", http.StatusConflict)
			return
		}
	}

	u.Balance -= e.Cost
	if result := tx.Save(&u); result.Error != nil {
		tx.Rollback()
		util.SendJSONError(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if err := tx.Model(&e).Association("Users").Append(&u); err != nil {
		tx.Rollback()
		util.SendJSONError(w, "Failed to add user to event", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit().Error; err != nil {
		util.SendJSONError(w, "Transaction commit error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "User joined event successfully",
		"event":   e.Name,
		"email":   u.Email,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Log the error of failing to encode or send the response
		fmt.Println("Error sending response:", err)
	}
}

//todo: implement
