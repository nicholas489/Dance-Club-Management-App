package main

import (
	"CPS406-Assignment-Backend/internal/api/http/server"
	"CPS406-Assignment-Backend/internal/db"
	"CPS406-Assignment-Backend/pkg/coach"
	"CPS406-Assignment-Backend/pkg/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func main() {
	// Connect to the database
	dataBase := db.ConnectDB()
	// Migrate the database
	db.MigrateDB(dataBase)

	dataBase.Create(&user.User{Name: "test1", Password: "test", Email: "test1@mail.com"})
	dataBase.Create(&user.User{Name: "test3", Password: "test", Email: "test2@mail.com"})
	dataBase.Create(&user.User{Name: "test4", Password: "test", Email: "test3@mail.com"})
	dataBase.Create(&coach.Coach{
		UserName:    "first",
		Email:       "first@gmail.com",
		PhoneNumber: 0,
		Password:    "test",
	})
	// Create a new router (chi router)

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// A good base middleware stack
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Routes for the API
	server.Server(r, dataBase)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("The server is running!"))
	})

	// Listen for requests on port 8080
	http.ListenAndServe(":8080", r)

}
