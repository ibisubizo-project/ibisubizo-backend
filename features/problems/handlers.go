package problems

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/render"
	"github.com/ofonimefrancis/problemsApp/features/users"
)

type MessageResponse struct {
	Message string `json:"error"`
}

func AddProblem(w http.ResponseWriter, r *http.Request) {
	var problem Problem

	err := json.NewDecoder(r.Body).Decode(&problem)
	if err != nil {
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error decoding payload"})
		return
	}

	if !bson.IsObjectIdHex(problem.CreatedBy) {
		log.Println("[AddProblem] Invalid User ID")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "User with specified ID doesn't exist."})
		return
	}

	if err := users.UserWithIDExists(problem.CreatedBy); err != nil {
		//How  are u even here, user with that ID doesn't exists
		log.Println(err)
		log.Println("[AddProblem] User with the ID specified doesn't exist")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "ID specified is not a valid UserID"})
		return
	}

	problem.ID = bson.NewObjectId()
	problem.CreatedAt = time.Now()

	if err := Create(problem); err != nil {
		log.Println(err)
		log.Println("[AddProblem] Error adding a new problem")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error creating a new problem"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, MessageResponse{Message: "Problem Successfully created"})

}

func GetApprovedPosts(w http.ResponseWriter, r *http.Request) {
	problems, err := ListAllApprovedListings()
	if err != nil {
		log.Println("[GetApprovedPosts] Error retrieving approved listings")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error retrieving posts"})
		return
	}
	if len(problems) == 0 {
		problems = []Problem{}
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, problems)
}

func GetAllListings(w http.ResponseWriter, r *http.Request) {
	problems, err := ListAll()
	if err != nil {
		log.Println("[GetAllListings] Error retrieving approved listings")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error retrieving posts"})
		return
	}
	if len(problems) == 0 {
		problems = []Problem{}
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, problems)
}

type ApproveRequest struct {
	ProblemID string `json:"id"`
}

func ApprovePost(w http.ResponseWriter, r *http.Request) {
	var request ApproveRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error decoding payload"})
		return
	}

	//Check if a problem with the ID Exists
	if !ProblemExists(request.ProblemID) {
		log.Println("There is no post with the specified ID")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "There is no post with the specified ID"})
		return
	}

	problem, err := GetByID(request.ProblemID)
	if err != nil {
		log.Println("Error retrieving post")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error retrieving post"})
		return
	}

	problem.IsApproved = true
	err = Update(request.ProblemID, problem)
	if err != nil {
		log.Println(err)
		log.Println("Error Approving post")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error approving post"})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, problem)

}
