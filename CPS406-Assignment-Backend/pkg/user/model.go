package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Password    string `json:"password"`
	Email       string `json:"email" gorm:"index,unique"`
	PhoneNumber int    `json:"phone_number"`
	Balance     int    `json:"balance"`
}
