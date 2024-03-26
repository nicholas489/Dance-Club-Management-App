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
	// Route for the login
	r.Route("/login", func(r chi.Router) {
		r.Post("/user", func(writer http.ResponseWriter, request *http.Request) {
			user.PostLogin(writer, request, db)
		})
		r.Post("/coach", func(writer http.ResponseWriter, request *http.Request) {
			coach.PostLogin(writer, request, db)
		})
	})
	// Route for the signup
	r.Route("/signup", func(r chi.Router) {
		r.Post("/user", func(writer http.ResponseWriter, request *http.Request) {
			user.PostSignup(writer, request, db)
		})
		r.Post("/coach", func(writer http.ResponseWriter, request *http.Request) {
			coach.PostSignup(writer, request, db)
		})
	})

	// Route for the user
	r.Route("/user", func(r chi.Router) {
		r.Use(util.JwtMiddlewareUser)
		r.Get("/{id}", func(writer http.ResponseWriter, request *http.Request) {
			user.GetUser(writer, request, db)
		})
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			user.GetAllUsers(writer, request, db)
		})
	})
	// Route for the coach
	// TODO: Implement the coach routes
	//r.Route("/coach", func(r chi.Router) {
	//	r.Use(util.CombinedJwtMiddleware(util.JwtMiddlewareCoach, util.JwtMiddlewareAdmin))
	//	r.Get("/{id}", func(writer http.ResponseWriter, request *http.Request) {
	//		coach.GetCoach(writer, request, db)
	//	})
	//	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
	//		coach.GetAllCoaches(writer, request, db)
	//
	//	})
	//})
	// Route for the event
	r.Route("/event", func(r chi.Router) {
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			coach.GetEvents(writer, request, db)
		})
		r.With(util.JwtMiddlewareCoach).Post("/", func(writer http.ResponseWriter, request *http.Request) {
			coach.PostEvent(writer, request, db)
		})

		r.Get("/{name}", func(writer http.ResponseWriter, request *http.Request) {
			coach.GetEvent(writer, request, db)
		})
		r.With(util.JwtMiddlewareUser).Post("/join", func(writer http.ResponseWriter, request *http.Request) {
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
