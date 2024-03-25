package event

import (
	"CPS406-Assignment-Backend/pkg/coach"
	"CPS406-Assignment-Backend/pkg/user"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name     string      `json:"name"`
	Couch    coach.Coach `json:"couch_id" gorm:"foreignKey:Email"`
	Location string      `json:"location"`
	Users    []user.User `json:"user_ids" gorm:"foreignKey:Email"`
}
