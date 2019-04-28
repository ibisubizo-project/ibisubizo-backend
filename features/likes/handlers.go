package likes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"github.com/ofonimefrancis/problemsApp/features/problems"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/render"
	"github.com/ofonimefrancis/problemsApp/features/comments"
)

type MessageResponse struct {
	Message string `json:"message,omitempty"`
}

type GenericPayload struct {
	ProblemID string `json:"problem_id,omitempty"`
	CommentID string `json:"comment_id,omitempty"`
	LikedBy   string `json:"liked_by"`
}

//AddLikeToComment - AddLikeToComment
func AddLikeToComment(w http.ResponseWriter, r *http.Request) {
	var request CommentLikes
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "Invalid Payload")
		return
	}
	request.ID = bson.NewObjectId()
	request.LikedOn = time.Now()

	if err := AddCommentLike(request); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error occurred while adding like"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, MessageResponse{Message: "Success"})
}

//AddLikeToProblem - AddLikeToProblem
func AddLikeToProblem(w http.ResponseWriter, r *http.Request) {
	var request GenericPayload
	var like ProblemLikes
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "Invalid Payload")
		return
	}
	like.ID = bson.NewObjectId()
	like.ProblemID = bson.ObjectIdHex(request.ProblemID)
	like.LikedBy = bson.ObjectIdHex(request.LikedBy)
	like.LikedOn = time.Now()

	if err := AddProblemLike(like); err != nil {
		log.Println(err)
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

//GetLikesForProblem - GetLikesForProblem
func GetLikesForProblem(w http.ResponseWriter, r *http.Request) {
	problemID := chi.URLParam(r, "problem_id")
	if !bson.IsObjectIdHex(problemID) {
		log.Println("Invalid bson hex object")
		return
	}

	if !problems.ProblemExists(problemID) {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Problem ID"})
		return
	}

	likes, err := GetAllLikesForProblem(problemID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error retrieving likes for post"})
		return
	}
	if len(likes) == 0 {
		likes = []ProblemLikes{}
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, likes)
}
