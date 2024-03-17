package server

import (
	"CPS406-Assignment-Backend/internal/api/http/user"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
)

func Server(r chi.Router, db *gorm.DB) {
	// Routes for the API
	r.Route("/users", func(r chi.Router) {
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			user.GetAllUsers(writer, request, db)
		})
	})
	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", func(writer http.ResponseWriter, request *http.Request) {
			user.GetUser(writer, request, db)
		})
		r.Post("/login", func(writer http.ResponseWriter, request *http.Request) {
			user.PostLogin(writer, request, db)
		})

	})

}
