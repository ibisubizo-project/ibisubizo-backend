package comments

import (
	"github.com/go-chi/chi"
)

//Routes - Routes for comments
func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/{postId}", AddComment)
	router.Get("/{postId}/public", GetCommentsForPublicPosts)
	router.Get("/{postId}/all", GetAllCommentsForPost)
	router.Get("/{postId}/{userId}/all", GetAllCommentsByUser)
	router.Put("/{commentId}/approve", ApproveComment)
	router.Get("/unapproved", GetAllUnapprovedComments)
	router.Delete("/{commentId}", RemoveComment)
	return router
}
