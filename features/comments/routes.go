package comments

import (
	"github.com/go-chi/chi"
)

//Routes - Routes for comments
func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/{postId}", AddComment)
	router.Get("/{postId}/all", GetCommentsForPost)
	return router
}
