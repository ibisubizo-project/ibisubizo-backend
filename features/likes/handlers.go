package likes

import (
	"encoding/json"
	"net/http"

	"github.com/ofonimefrancis/problemsApp/features/problems"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/render"
	"github.com/ofonimefrancis/problemsApp/features/comments"
)

type MessageResponse struct {
	Message string `json:"message,omitempty"`
}

//AddLikeToComment - AddLikeToComment
func AddLikeToComment(w http.ResponseWriter, r *http.Request) {
	var request Likes
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "Invalid Payload")
		return
	}
	request.ID = bson.NewObjectId()
	request.ProblemID = ""

	if err := AddLike(request); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error occurred while adding like"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, MessageResponse{Message: "Success"})
}

//AddLikeToProblem - AddLikeToProblem
func AddLikeToProblem(w http.ResponseWriter, r *http.Request) {
	var request Likes
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "Invalid Payload")
		return
	}
	request.ID = bson.NewObjectId()
	request.CommentID = ""

	if err := AddLike(request); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error occurred while adding like"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, MessageResponse{Message: "Success"})
}

//Payload - Request Payload
type Payload struct {
	ID        string `json:"id"`
	CommentID string `json:"comment_id"`
	ProblemID string `json:"problem_id"`
}

//DeleteLikeFromComment - DeleteLikeFromComment
func DeleteLikeFromComment(w http.ResponseWriter, r *http.Request) {
	//We need the like ID and comment ID
	var like Payload
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Payload"})
		return
	}

	if !comments.CommentExists(like.CommentID) {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Comment ID"})
		return
	}

	if err := DeleteLikeForComment(like.ID, like.CommentID); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error removing like for comment"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, MessageResponse{Message: "Like Removed"})
}

//DeleteLikeFromProblem - DeleteLikeFromProblem
func DeleteLikeFromProblem(w http.ResponseWriter, r *http.Request) {
	var like Payload
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Payload"})
		return
	}
	if !problems.ProblemExists(like.ProblemID) {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Problem ID"})
		return
	}

	if err := DeleteLikeForProblen(like.ID, like.ProblemID); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error removing like for problem"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, MessageResponse{Message: "Like Removed"})
}
