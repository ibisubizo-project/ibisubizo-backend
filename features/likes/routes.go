package likes

import (
	"github.com/go-chi/chi"
)

//Route - Route
func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/add/to/comment", AddLikeToComment)
	router.Post("/add/to/problem", AddLikeToProblem)
	router.Get("/{problem_id}", GetLikesForProblem)
	router.Delete("/remove/from/comment", DeleteLikeFromComment)
	router.Delete("/remove/from/problem", DeleteLikeFromProblem)

	return router
}
