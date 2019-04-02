package users

import (
	"github.com/go-chi/chi"
)

const (
	//Register - Register route
	RegisterRoute = "/register"
	LoginRoute    = "/login"
)

//Routes -  All the user specific routes
func AuthRoutes() *chi.Mux {
	router := chi.NewMux()
	router.Post(RegisterRoute, RegisterUser)
	router.Post(LoginRoute, Login)
	return router
}
