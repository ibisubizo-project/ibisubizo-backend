package users

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/ofonimefrancis/problemsApp/config"
)

const (
	//RegisterRoute - Register route
	RegisterRoute = "/register"
	LoginRoute    = "/login"
)

//Routes -  All the user specific routes
func AuthRoutes() *chi.Mux {
	router := chi.NewMux()

	router.Post(RegisterRoute, RegisterUser)

	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.GetTokenAuth()))

		r.Use(jwtauth.Authenticator)

		router.Post(LoginRoute, Login)
	})

	return router
}
