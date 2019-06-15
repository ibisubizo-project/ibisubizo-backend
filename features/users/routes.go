package users

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/ofonimefrancis/problemsApp/config"
)

const (
	//RegisterRoute - Register route
	RegisterRoute   = "/register"
	LoginRoute      = "/login"
	AdminLoginRoute = "/admin/login"
	UserRoute       = "/{user_id}"
	AllUsersRoute   = "/"
)

//AuthRoutes -  All the Authentication specific routes
func AuthRoutes() *chi.Mux {
	router := chi.NewMux()

	router.Post(RegisterRoute, RegisterUser)

	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.GetTokenAuth()))

		r.Use(jwtauth.Authenticator)

		router.Post(LoginRoute, Login)
		router.Post(AdminLoginRoute, AdminLogin)
	})

	return router
}

//UserRoutes - UserRoutes
func UserRoutes() *chi.Mux {
	router := chi.NewMux()
	router.Get(AllUsersRoute, RetrieveAllUsers)
	router.Get(UserRoute, FetchUserByID)

	return router
}
