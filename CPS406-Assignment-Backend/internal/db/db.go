package db

import (
	"CPS406-Assignment-Backend/pkg/coach"
	"CPS406-Assignment-Backend/pkg/event"
	"CPS406-Assignment-Backend/pkg/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDB is a function that connects to the database
func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("recreation.db"), &gorm.Config{})
	print("Connected to the database	\n")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// MigrateDB is a function that creates the tables in the database
func MigrateDB(db *gorm.DB) {
	print("Migrating the database \n")
	db.AutoMigrate(&user.User{}, &coach.Coach{}, &event.Event{})
	print("Database migrated \n")
}
