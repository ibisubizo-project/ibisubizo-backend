package comments

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/problemsApp/features/problems"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

//MessageResponse - MessageResponse
type MessageResponse struct {
	Message string `json:"error"`
}

//GetCommentsForPost - GetCommentsForPost
func GetCommentsForPost(w http.ResponseWriter, r *http.Request) {
	postID := strings.TrimSpace(chi.URLParam(r, "postId"))

	//Check if a problem with postID exists
	ok := problems.ProblemExists(postID)
	if !ok {
		log.Println("[GetCommentsForPost] Problem with the specified ID does not exists")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Problem ID"})
		return
	}

	comments, err := GetCommentsForProblem(postID)
	if err != nil {
		log.Println("[GetCommentsForPost] Error retrieving comments for problem")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error retrieving comments"})
		return
	}

	if len(comments) == 0 {
		comments = []Comment{}
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, comments)
}

//AddComment - AddComment
func AddComment(w http.ResponseWriter, r *http.Request) {
	var comment Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Println("[AddComment] Error decoding payload")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Payload"})
		return
	}
	comment.ID = bson.NewObjectId()
	comment.CommentedAt = time.Now()

	if err := CreateComment(comment); err != nil {
		log.Println("[AddComment] Error adding a new comment")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Unable to add comment"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, MessageResponse{Message: "Comment added"})

}
