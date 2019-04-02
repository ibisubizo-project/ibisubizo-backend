package users

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/render"
	"github.com/ofonimefrancis/problemsApp/config"
)

//ErrorResponse - ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}

//LoginSuccessResponse - LoginSuccessResponse
type LoginSuccessResponse struct {
	User        Users  `json:"user"`
	TokenString string `json:"token"`
}

//RegisterUser - Signs up a user to the platform
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Something is not right..."})
		return
	}

	if user.Exists() {
		log.Println("[RegisterUser] User exists..")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "User with phone already exists"})
		return
	}

	passwordHash, err := NewPasswordHash(user.Password)
	if err != nil {
		log.Println("[RegisterUser] Error creating password hash")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Error generating password hash"})
		return
	}

	user.HashedPassword = passwordHash.Hash
	user.Salt = passwordHash.Salt

	//Assign ID
	user.ID = bson.NewObjectId()
	user.IsAdmin = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := Create(user); err != nil {
		log.Println(err)
		log.Println("[RegisterUser] Error creating a new user")

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Error creating user"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

//LoginObject - LoginObject
type LoginObject struct {
	PhoneNumber string `json:"phone"`
	Password    string `json:"password"`
}

//Login - Login
func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginObject

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Println("[Login] Error decoding payload")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Something went wrong..."})
		return
	}

	if !UserExists(loginRequest.PhoneNumber) {
		log.Println("[Login] User doesn't exist")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Invalid Phone number or password"})
		return
	}

	user, err := Read(loginRequest.PhoneNumber)
	if err != nil {
		log.Println("[Login] Error retrieving user")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Something went wrong..."})
		return
	}

	if !VerifyPassword(loginRequest.Password, user.HashedPassword, user.Salt) {
		//Invalid Password
		log.Println("[Login] Invalid Password")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Invalid PhoneNumber or Password"})
		return
	}

	jwtAuth := config.GetTokenAuth()
	_, tokenString, err := jwtAuth.Encode(jwt.MapClaims{
		"id":        user.ID,
		"isAdmin":   user.IsAdmin,
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"phone":     user.PhoneNumber,
	})
	if err != nil {
		log.Println("[Login] Error Encoding JWT payload")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Something went wrong..."})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, LoginSuccessResponse{User: user, TokenString: tokenString})
}
