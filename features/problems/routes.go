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

	return router
}
