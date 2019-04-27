package problems

import (
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/", AddProblem)
	router.Get("/", GetAllListings)
	router.Get("/approved", GetApprovedPosts)
	router.Post("/approve", ApprovePost)
	router.Post("/resolve", ResolveProblem)
	router.Get("/user/{user_id}", GetUserProblems)
	router.Get("/{problem_id}", GetProblem)

	return router
}
