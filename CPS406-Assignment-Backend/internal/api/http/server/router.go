package server

import (
	"CPS406-Assignment-Backend/internal/api/http/coach"
	"CPS406-Assignment-Backend/internal/api/http/user"
	"CPS406-Assignment-Backend/internal/util"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
)

func Server(r chi.Router, db *gorm.DB) {
	// Routes for the API

	r.Route("/login", func(r chi.Router) {
		r.Post("/user", func(writer http.ResponseWriter, request *http.Request) {
			user.PostLogin(writer, request, db)
		})
		r.Post("/coach", func(writer http.ResponseWriter, request *http.Request) {
			coach.PostLogin(writer, request, db)
		})
	})
	r.Route("/signup", func(r chi.Router) {
		r.Post("/user", func(writer http.ResponseWriter, request *http.Request) {
			user.PostSignup(writer, request, db)
		})
		r.Post("/coach", func(writer http.ResponseWriter, request *http.Request) {
			coach.PostSignup(writer, request, db)
		})
	})
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
		r.Post("/join/event", func(writer http.ResponseWriter, request *http.Request) {
			user.JoinEvent(writer, request, db)
		})

	})

	r.Route("/coach", func(r chi.Router) {
		r.Use(util.CombinedJwtMiddleware(util.JwtMiddlewareCoach, util.JwtMiddlewareAdmin))
		r.Post("/event/make", func(writer http.ResponseWriter, request *http.Request) {
			coach.PostEvent(writer, request, db)
		})
		r.Get("/{name}", func(writer http.ResponseWriter, request *http.Request) {
			coach.GetEvent(writer, request, db)
		})
		r.Delete("/delete/{email}", func(writer http.ResponseWriter, request *http.Request) {
			coach.DeleteUser(writer, request, db)
		})

	})

}
