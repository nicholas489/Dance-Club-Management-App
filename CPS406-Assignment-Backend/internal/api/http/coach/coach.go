package coach

import (
	"CPS406-Assignment-Backend/internal/util"
	"CPS406-Assignment-Backend/pkg/coach"
	"CPS406-Assignment-Backend/pkg/event"
	"CPS406-Assignment-Backend/pkg/jwtM"
	"CPS406-Assignment-Backend/pkg/user"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
)

func PostEvent(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Parse the request body
	var event, existingEvent event.Event
	err := json.NewDecoder(request.Body).Decode(&event)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// If the event already exists
	result := db.First(&existingEvent, "name = ?", event.Name)
	if result.Error == nil {
		util.SendJSONError(writer, "Event already exists", http.StatusConflict)
		return
	}
	// make event and putting an empty user list
	event.Users = []user.User{}
	// Create the event in the database
	db.Create(&event)
	// Send the event as a response and set the status code to 201 (Created)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(event)

}

func GetEvent(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Get the id from the url
	name := chi.URLParam(request, "name")

	// Get the event from the database and save it in the event variable
	var event event.Event
	db.First(&event, "Name = ?", name)

	// Send the event as a response
	json.NewEncoder(writer).Encode(event)
}

func GetEvents(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Get all the events from the database
	var events []event.Event
	db.Find(&events)

	// Send the events as a response
	json.NewEncoder(writer).Encode(events)
}

func PostLogin(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Parse the request body
	var login coach.Coach
	err := json.NewDecoder(request.Body).Decode(&login)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// Check if the coach exists in the database
	var coach coach.Coach
	result := db.First(&coach, "email = ?", login.Email)
	if result.Error != nil {
		util.SendJSONError(writer, "Coach not found", http.StatusNotFound)
		return
	}
	// If the coach exists, check if the passwords match
	if coach.Password != login.Password {
		util.SendJSONError(writer, "Invalid password", http.StatusUnauthorized)
		return
	}
	privileges := util.SetPrivileges(jwtM.CustomClaims{Privileges: jwtM.Privileges{Coach: true}})
	// Generate a JWT token
	token, err := util.GenerateJWT(coach.Email, privileges)
	if err != nil {
		util.SendJSONError(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	// Send the token as a response
	writer.Header().Set("Authorization", "Bearer "+token)
	writer.WriteHeader(http.StatusOK)
	// Send the coach as a response
	json.NewEncoder(writer).Encode(coach)
}

func PostSignup(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Parse the request body
	var coach coach.Coach
	err := json.NewDecoder(request.Body).Decode(&coach)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// If the coach already exists
	result := db.First(&coach, "email = ?", coach.Email)
	if result.Error == nil {
		util.SendJSONError(writer, "Coach already exists", http.StatusConflict)
		return
	}
	// Generate a JWT token
	privileges := util.SetPrivileges(jwtM.CustomClaims{Privileges: jwtM.Privileges{Coach: true}})
	token, err := util.GenerateJWT(coach.Email, privileges)
	// Create the coach in the database
	db.Create(&coach)
	// Send the coach as a response and set the status code to 201 (Created)
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Authorization", "Bearer "+token)
	json.NewEncoder(writer).Encode(coach)
}
