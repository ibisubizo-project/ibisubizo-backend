package problems

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ofonimefrancis/problemsApp/features/users"
)

//MessageResponse - MessageResponse
type MessageResponse struct {
	Message string `json:"error"`
}

//ApproveRequest - ApproveRequest
type ApproveRequest struct {
	ProblemID string `json:"id"`
}

//AddProblem - AddProblem
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
	problem.IsResolved = false

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

//GetAllResolvedPosts - Get all Resolved problems
func GetAllResolvedPosts(w http.ResponseWriter, r *http.Request) {
	problems, err := ListAllResolvedProblems(0)
	if err != nil {
		log.Println(err)
		log.Println("[GetAllResolvedPosts] Error retrieving resolved problems")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error retrieving resolved problems"})
		return
	}
	if len(problems) == 0 {
		problems = []Problem{}
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, problems)
}

//GetAllUnResolvedPosts - Fetch UnResolved problems
func GetAllUnResolvedPosts(w http.ResponseWriter, r *http.Request) {
	problems, err := ListAllUnResolvedProblems(0)
	if err != nil {
		log.Println(err)
		log.Println("[GetAllUnResolvedPosts] Error retrieving unresolved problems")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error retrieving unresolved problems"})
		return
	}
	if len(problems) == 0 {
		problems = []Problem{}
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, problems)
}

//GetApprovedPosts - GetApprovedPosts
func GetApprovedPosts(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page := 1
	var err error

	if pageStr == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
	}
	problems, err := ListAllApprovedListings(page)
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

//GetAllListings - GetAllListings
func GetAllListings(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page := 1
	var err error

	if pageStr == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
	}

	problems, err := ListAll(page)
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

//ApprovePost - ApprovePost
func ApprovePost(w http.ResponseWriter, r *http.Request) {
	var request ApproveRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		log.Println("[ApprovePost] Error decoding payload")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid payload"})
		return
	}

	//Check if a problem with the ID Exists
	if !ProblemExists(request.ProblemID) {
		log.Println("[ApprovePost] There is no post with the specified ID")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "There is no post with the specified ID"})
		return
	}

	problem, err := GetByID(request.ProblemID)
	if err != nil {
		log.Println("[ApprovePost] Error retrieving post")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error retrieving post"})
		return
	}

	problem.IsApproved = true
	err = Update(request.ProblemID, problem)
	if err != nil {
		log.Println(err)
		log.Println("[ApprovePost] Error Approving post")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error approving post"})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, problem)

}

//ResolveProblem - ResolveProblem
func ResolveProblem(w http.ResponseWriter, r *http.Request) {
	var request ApproveRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println(err)
		log.Println("[ResolveProblem] Error decoding payload")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid payload"})
		return
	}

	//Check if a problem with the ID Exists
	if !ProblemExists(request.ProblemID) {
		log.Println("[ResolveProblem] There is no post with the specified ID")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "There is no post with the specified ID"})
		return
	}

	problem, err := GetByID(request.ProblemID)
	if err != nil {
		log.Println(err)
		log.Println("[ResolveProblem] Error retrieving problem with the specified ID")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error retrieving problem"})
		return
	}

	problem.IsResolved = true
	err = Update(request.ProblemID, problem)
	if err != nil {
		log.Println(err)
		log.Println("[ResolveProblem] Error Approving post")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error approving post"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, MessageResponse{Message: "Problem resolved"})
}

//GetUserProblems - GetUserProblems
func GetUserProblems(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	pageStr := r.URL.Query().Get("page")
	page := 1
	var err error

	if pageStr == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
	}
	problems, err := GetUserListings(userID, page)
	if err != nil {
		log.Println(err)
		log.Println("[GetUserProblems] Error Retrieving users problem")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Something went wrong..."})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, problems)
}

func GetProblem(w http.ResponseWriter, r *http.Request) {
	problemID := chi.URLParam(r, "problem_id")

	ok := ProblemExists(problemID)
	if !ok {
		log.Println("Problem with that ID doesn't exists")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Problem ID"})
		return
	}
	problem, err := GetByID(problemID)
	if err != nil {
		log.Println("Error retrieving problem")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error retrieving problem"})
		return
	}

	// var result ProblemDetail

	// allComments, err := comments.GetCommentsForProblem(problem.ID)
	// if err != nil {
	// 	log.Println(err)
	// 	log.Println("Error retrieving comments")
	// 	render.Status(r, http.StatusBadRequest)
	// 	render.JSON(w, r, MessageResponse{Message: "Error retrieving comments"})
	// 	return
	// }

	// result.Problem = problem
	// result.Comments = allComments

	render.Status(r, http.StatusOK)
	render.JSON(w, r, problem)
}
