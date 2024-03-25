package db

import (
	"CPS406-Assignment-Backend/pkg/coach"
	"CPS406-Assignment-Backend/pkg/event"
	"CPS406-Assignment-Backend/pkg/user"
	"gorm.io/gorm"
	"log"
)

func SeedDatabase(db *gorm.DB) {
	// Sample coaches
	coaches := []coach.Coach{
		{Name: "CoachMike", Email: "mike@example.com", PhoneNumber: 1234567890, Password: "pass123"},
		{Name: "CoachAnna", Email: "anna@example.com", PhoneNumber: 1234567891, Password: "pass456"},
	}

	// Create coaches
	for _, coach := range coaches {
		result := db.Create(&coach)
		if result.Error != nil {
			log.Printf("Could not create coach: %v", result.Error)
			continue
		}
	}

	// Sample event
	events := []event.Event{
		{Name: "Morning Yoga", CoachEmail: "mike@example.com", Location: "Central Park", Cost: 10},
		{Name: "Evening Run", CoachEmail: "anna@example.com", Location: "Riverside", Cost: 5},
	}

	// Create events and associate with coaches
	for _, event := range events {
		// Assuming the Coach is identified by their email address.
		var coach coach.Coach
		if err := db.Where("email = ?", event.CoachEmail).First(&coach).Error; err != nil {
			log.Printf("Coach not found: %v", err)
			continue
		}
		event.CoachEmail = coach.Email

		result := db.Create(&event)
		if result.Error != nil {
			log.Printf("Could not create event: %v", result.Error)
			continue
		}
	}

	// Sample users
	users := []user.User{
		{Name: "John Doe", Email: "john.doe@example.com", PhoneNumber: 1234567892, Password: "secure123", Balance: 100},
		{Name: "Jane Smith", Email: "jane.smith@example.com", PhoneNumber: 1234567893, Password: "secure456", Balance: 150},
	}

	// Create users and enroll to events
	for i, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			log.Printf("Could not create user: %v", result.Error)
			continue
		}

	}
}
