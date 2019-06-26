package users

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/ofonimefrancis/problemsApp/config"
)

const (
	//RegisterRoute - Register route
	RegisterRoute       = "/register"
	LoginRoute          = "/login"
	AdminLoginRoute     = "/admin/login"
	AdmiCreationRoute   = "/admin/register"
	UserRoute           = "/{user_id}"
	ForgetPasswordRoute = "/forget"
	AllUsersRoute       = "/"
	ConfirmRestRoute    = "/confirmation/{reset_token}"
	ChangePasswordRoute = "/changepassword"
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
		router.Post(AdmiCreationRoute, CreateAdmin)
	})

	return router
}

//UserRoutes - UserRoutes
func UserRoutes() *chi.Mux {
	router := chi.NewMux()
	router.Get(AllUsersRoute, RetrieveAllUsers)
	router.Get(UserRoute, FetchUserByID)
	router.Post(ForgetPasswordRoute, ForgetPassword)
	router.Post(ConfirmRestRoute, ConfirmResetToken)
	router.Post(ChangePasswordRoute, ChangePassword)

	return router
}
