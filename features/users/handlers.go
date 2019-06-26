package users

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/render"
	"github.com/ofonimefrancis/problemsApp/config"
	"github.com/ofonimefrancis/problemsApp/utils"
)

//ErrorResponse - ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}

//SuccessResponse - LoginSuccessResponse
type SuccessResponse struct {
	User        Users  `json:"user"`
	TokenString string `json:"token"`
}

//RegisterUser - Signs up a user to the platform
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println()
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

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("[RegisterUser] Error creating password hash")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: err.Error()})
		return
	}

	// passwordHash, err := NewPasswordHash(user.Password)
	// if err != nil {
	// 	log.Println("[RegisterUser] Error creating password hash")
	// 	log.Println(err)
	// 	render.Status(r, http.StatusBadRequest)
	// 	render.JSON(w, r, ErrorResponse{Error: err.Error()})
	// 	return
	// }

	user.Password = "" //Dont disclose user password
	user.HashedPassword = passwordHash
	//user.Salt = passwordHash.Salt

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

	_, tokenString, err := config.GetTokenAuth().Encode(jwt.MapClaims{
		"id":        user.ID,
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"phone":     user.PhoneNumber,
	})
	if err != nil {
		log.Println("[RegisterUser] Error encoding jwt payload")
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, ErrorResponse{Error: "Error encoding jwt payload"})
		return
	}

	expires := time.Now().AddDate(1, 0, 0)
	ck := http.Cookie{
		Name:     "jwt",
		HttpOnly: false,
		Path:     "/",
		Expires:  expires,
		Value:    tokenString,
	}

	// write the cookie to response
	http.SetCookie(w, &ck)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, SuccessResponse{User: user, TokenString: tokenString})
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

	//fmt.Printf("Password: %s\nHash: %s\nSalt: %s\n", loginRequest.Password, user.HashedPassword, user.Salt)

	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(loginRequest.Password)); err != nil {
		//Invalid Password
		log.Println("[Login] Invalid Password")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Invalid PhoneNumber or Password"})
		return
	}

	// if !VerifyPassword(loginRequest.Password, user.HashedPassword, user.Salt) {
	// 	//Invalid Password
	// 	log.Println("[Login] Invalid Password")
	// 	render.Status(r, http.StatusBadRequest)
	// 	render.JSON(w, r, ErrorResponse{Error: "Invalid PhoneNumber or Password"})
	// 	return
	// }

	jwtAuth := config.GetTokenAuth()
	_, tokenString, err := jwtAuth.Encode(jwt.MapClaims{
		"id":        user.ID,
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

	expires := time.Now().AddDate(1, 0, 0)
	ck := http.Cookie{
		Name:     "jwt",
		HttpOnly: false,
		Path:     "/",
		Expires:  expires,
		Value:    tokenString,
	}

	// write the cookie to response
	http.SetCookie(w, &ck)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, SuccessResponse{User: user, TokenString: tokenString})
}

//AdminLogin - Login an admin user
func AdminLogin(w http.ResponseWriter, r *http.Request) {
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
	if !user.IsAdmin {
		log.Println("User is not an admin")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Please Login as an admin"})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(loginRequest.Password)); err != nil {
		//Invalid Password
		log.Println("[Login] Invalid Password")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Invalid PhoneNumber or Password"})
		return
	}

	jwtAuth := config.GetTokenAuth()
	_, tokenString, err := jwtAuth.Encode(jwt.MapClaims{
		"id":        user.ID,
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

	expires := time.Now().AddDate(1, 0, 0)
	ck := http.Cookie{
		Name:     "jwt",
		HttpOnly: false,
		Path:     "/",
		Expires:  expires,
		Value:    tokenString,
	}

	// write the cookie to response
	http.SetCookie(w, &ck)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, SuccessResponse{User: user, TokenString: tokenString})
}

//FetchUserByID - FetchUserByID
func FetchUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	if len(userID) == 0 {
		log.Println("Invalid User ID")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Invalid User ID"})
		return
	}

	user, err := GetUserById(userID)
	if err != nil {
		log.Println("[FetchUserById] Error retrieving user by ID")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Error retrieving user with specified ID"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

func RetrieveAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := ReadAll()
	if err != nil {
		log.Println("[RetrieveAllUsers] Error retrieving all users")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Something bad happened"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, users)
}
