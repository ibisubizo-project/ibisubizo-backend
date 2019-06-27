package users

import (
	"encoding/json"
	"fmt"
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

type ChangePasswordRequest struct {
	Password string `json:"password"`
	UserID   string `json:"user_id"`
}

//SuccessResponse - LoginSuccessResponse
type SuccessResponse struct {
	User        Users  `json:"user"`
	TokenString string `json:"token"`
}

//CreateAdmin - Creates an Admin User
func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var user Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Something is not right..."})
		return
	}

	if user.Exists() {
		log.Println("[CreateAdmin] User exists..")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "User with phone already exists"})
		return
	}

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("[CreateAdmin] Error creating password hash")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: err.Error()})
		return
	}

	user.Password = "" //Dont disclose user password
	user.HashedPassword = passwordHash

	//Assign ID
	user.ID = bson.NewObjectId()
	user.IsAdmin = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := Create(user); err != nil {
		log.Println(err)
		log.Println("[CreateAdmin] Error creating a new user")

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
		log.Println("[CreateAdmin] Error encoding jwt payload")
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

	user.Password = "" //Dont disclose user password
	user.HashedPassword = passwordHash

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

	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(loginRequest.Password)); err != nil {
		//Invalid Password
		log.Println("[Login] Invalid Password")
		log.Println("[Bcrypt]")
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
		log.Println("[Bcrypt]")
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

//ForgetPassword - First Stage of forget password
func ForgetPassword(w http.ResponseWriter, r *http.Request) {
	user := Users{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("[ForgetPassword] Error decoding payload")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Invalid Payload"})
		return
	}
	//get the phone in the request body
	//Check if user with that phone exists
	if !UserExists(user.PhoneNumber) {
		log.Println("[ForgetPassword] Error decoding payload")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "User Not Found"})
		return
	}
	verifiedUser, err := Read(user.PhoneNumber)
	if err != nil {
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "User Not Found"})
		return
	}
	user = verifiedUser

	randomDigits, _ := utils.GenerateRandomString(5)
	randomToken, _ := utils.GenerateRandomString(35)
	user.ResetCode = randomDigits
	user.ResetToken = randomToken

	user.UpdatedAt = time.Now()
	fiveMinutes := time.Minute * time.Duration(5)
	expiresIn := time.Now().Add(fiveMinutes)

	fmt.Println("In five minutes will be ", expiresIn)
	user.CodeExpiresAt = expiresIn

	log.Println(user)

	log.Println("ObjectID ", verifiedUser.ID)
	err = Update(verifiedUser, user)
	if err != nil {
		log.Println("[ForgetPassword] Error Updating User Information")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Something Bad happened"})
		return
	}

	//Send Reset Code and link to reset password
	resetTokenURL := fmt.Sprintf("%s/auth/confirm/token/%s", config.APP_URL, user.ResetToken)
	var message = fmt.Sprintf("Your reset code is %s\nVisit %s and enter the above code", user.ResetCode, resetTokenURL)
	utils.SendSMS(verifiedUser.PhoneNumber, message, "Ibisubizo")
	render.Status(r, http.StatusOK)
	render.JSON(w, r, ErrorResponse{Error: "A Reset Code valid for 5 mins has been sent to your phone"})
}

func ConfirmResetToken(w http.ResponseWriter, r *http.Request) {
	//https://api.ibisubizo.com/api/auth/{reset_token}
	//RequestBody { 'reset_code'}
	//Check if the reset_token and reset_code match a user
	//Check if the reset_token and code hasn't expired
	//If expired, tell them to resend the reset token (Create Reset token and code and adds it to the database)
	//ElseReturn OK nd user can proceed to the next step

	type ResetCodeParams struct {
		ResetCode string `json:"reset_code"`
	}

	resetToken := chi.URLParam(r, "reset_token")
	var resetCodeBody ResetCodeParams
	err := json.NewDecoder(r.Body).Decode(&resetCodeBody)
	if err != nil {
		log.Println(err)
		log.Println("[ConfirmResetToken] Invalid Payload")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Invalid Payload"})
		return
	}

	user, err := ConfirmResetTokens(resetToken, resetCodeBody.ResetCode)

	if err != nil {
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: err.Error()})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, struct {
		Message string `json:"message"`
		User    Users  `json:"user"`
	}{Message: "Confirmation Successful", User: user})
}

//ChangePassword - ChangePassword
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var changePasswordBody ChangePasswordRequest
	token := r.URL.Query().Get("token")
	err := json.NewDecoder(r.Body).Decode(&changePasswordBody)
	if err != nil {
		log.Println(err)
		log.Println("[ChangePassword] Invalid Payload")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Invalid Payload"})
		return
	}

	//Check If user with reset token and ID exists
	if !UserWithTokenExists(changePasswordBody.UserID, token) {
		log.Println(err)
		log.Println("[ChangePassword] Invalid Token")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Invalid Token"})
		return
	}

	user, err := GetUserById(changePasswordBody.UserID)
	if err != nil {
		log.Println(err)
		log.Println("[ChangePassword] Unable to retrieve User with specified ID")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "User with ID not found"})
		return
	}

	passwordHash, err := utils.HashPassword(changePasswordBody.Password)
	if err != nil {
		log.Println(err)
		log.Println("[ChangePassword] Error while hashing password")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Something went wrong.."})
		return
	}
	user.HashedPassword = passwordHash
	user.UpdatedAt = time.Now()
	user.ResetCode = ""
	user.ResetToken = ""
	user.CodeExpiresAt = time.Time{}

	if err = UpdateByID(user.ID.Hex(), user); err != nil {
		log.Println(err)
		log.Println("[ChangePassword] Error while updating user password")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Error: "Something went wrong.."})
		return
	}
	render.Status(r, http.StatusOK)
	//Send SMS, your new password has been updated
	render.JSON(w, r, ErrorResponse{Error: "Password Successfully changed..."})
	return

}
