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
	Message string `json:"message"`
}

type NewCommentResponse struct {
	Message string  `json:"message"`
	Comment Comment `json:"comment,omitempty"`
}

//GetCommentsForPublicPosts - GetCommentsForPublicPosts
func GetCommentsForPublicPosts(w http.ResponseWriter, r *http.Request) {
	postID := strings.TrimSpace(chi.URLParam(r, "postId"))

	//Check if a problem with postID exists
	ok := problems.ProblemExists(postID)
	if !ok {
		log.Println("[GetCommentsForPost] Problem with the specified ID does not exists")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Problem ID"})
		return
	}
	//Check If the post is (public)
	problemStatus := problems.GetProblemStatus(postID)
	if problemStatus != "" && problemStatus != "private" {
		log.Println("Status here should be public: ", problemStatus)
		comments, err := GetCommentsForProblem(postID)
		if err != nil {
			log.Println("[GetCommentsForPost] Error retrieving comments for problem")
			log.Println(err)
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, MessageResponse{Message: "Error retrieving comments"})
			return
		}
		approvedComments := []Comment{}
		for _, comment := range comments {
			if !comment.IsApproved {
				continue
			}
			approvedComments = append(approvedComments, comment)
		}

		if len(approvedComments) == 0 {
			approvedComments = []Comment{}
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, approvedComments)
		return
	}
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, MessageResponse{Message: "Private Comments will not be displayed"})
	return

}

//GetAllCommentsForPost - GetAllCommentsForPost
func GetAllCommentsForPost(w http.ResponseWriter, r *http.Request) {
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
	return
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
	comment.IsApproved = false

	if err := CreateComment(comment); err != nil {
		log.Println("[AddComment] Error adding a new comment")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Unable to add comment"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, NewCommentResponse{Message: "Comment added", Comment: comment})

}

//ApproveComment - ApproveComment
func ApproveComment(w http.ResponseWriter, r *http.Request) {
	commentID := chi.URLParam(r, "commentId")

	if CommentExists(commentID) {
		comment, _ := GetComment(commentID)
		comment.IsApproved = true
		err := Update(commentID, comment)
		if err != nil {
			log.Println("Error Updating Comment")
			log.Println(err)
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, MessageResponse{Message: "Error Updating Comment"})
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, MessageResponse{Message: "Comment Approved for public view"})
		return
	}
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, MessageResponse{Message: "Comment with the specified ID doesn't exists"})
	return
}

//GetAllUnapprovedComments - GetAllUnapprovedComments
func GetAllUnapprovedComments(w http.ResponseWriter, r *http.Request) {
	comments, err := GetAllUnapproved()
	if err != nil {
		log.Println(err)
		log.Println("Error retrieving all unapproved comments")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Unable to retrieve unapproved comments"})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, comments)
	return
}

//RemoveComment - RemoveComment
func RemoveComment(w http.ResponseWriter, r *http.Request) {
	commentID := chi.URLParam(r, "commentId")
	err := Remove(commentID)
	if err != nil {
		log.Println("Error removing comment")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error Removing Comment."})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, MessageResponse{Message: "Comment deleted successfully."})
	return

}
