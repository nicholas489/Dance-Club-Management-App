package coach

import (
	"CPS406-Assignment-Backend/internal/util"
	"CPS406-Assignment-Backend/pkg/coach"
	"CPS406-Assignment-Backend/pkg/event"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
)

func PostCoach(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
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
	// Create the coach in the database
	db.Create(&coach)
	// Send the coach as a response and set the status code to 201 (Created)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(coach)

}

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
	//
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
