package event

import (
	"CPS406-Assignment-Backend/pkg/user"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name       string      `json:"name" gorm:"index;unique"`
	CoachEmail string      `json:"coach_email"` // Store email to fetch and associate Coach
	Location   string      `json:"location"`
	Cost       int         `json:"cost"`
	Users      []user.User `json:"users"`
}

type UserEvent struct {
	user.User
	Event
}
