package server

import (
	"CPS406-Assignment-Backend/internal/api/http/user"
	"CPS406-Assignment-Backend/internal/util"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
)

func Server(r chi.Router, db *gorm.DB) {
	// Routes for the API
	r.Route("/users", func(r chi.Router) {
		r.Use(util.CombinedJwtMiddleware(util.JwtMiddlewareCoach, util.JwtMiddlewareAdmin))
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			user.GetAllUsers(writer, request, db)
		})
	})
	r.Route("/user", func(r chi.Router) {
		r.Use(util.JwtMiddlewareUser)
		r.Get("/{id}", func(writer http.ResponseWriter, request *http.Request) {
			user.GetUser(writer, request, db)
		})

	})
	r.Route("/login", func(r chi.Router) {
		r.Post("/", func(writer http.ResponseWriter, request *http.Request) {
			user.PostLogin(writer, request, db)
		})
	})
	r.Route("/signup", func(r chi.Router) {
		r.Post("/", func(writer http.ResponseWriter, request *http.Request) {
			user.PostSignup(writer, request, db)
		})
	})

}
