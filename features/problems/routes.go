package problems

import (
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/", AddProblem)
	router.Get("/", GetAllListings)
	router.Get("/approved", GetApprovedPosts)
	router.Get("/resolved", GetAllResolvedPosts)
	router.Get("/unresolved", GetAllUnResolvedPosts)
	router.Put("/approve", ApprovePost)
	router.Get("/unapproved", UnApprovedPost)
	router.Put("/resolve", ResolveProblem)
	router.Get("/user/{user_id}", GetUserProblems)
	router.Get("/{problem_id}", GetProblem)
	router.Put("/{problem_id}/{user_id}", UpdateMyPost)
	router.Delete("/{problem_id}", DeleteMyPost)

	return router
}
