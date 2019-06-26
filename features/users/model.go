package users

import (
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/problemsApp/config"
)

//Users - Represents a user in the system
type Users struct {
	ID             bson.ObjectId `json:"_id,omitempty"`
	FirstName      string        `json:"firstname"`
	MiddleName     string        `json:"middlename"`
	LastName       string        `json:"lastname"`
	PhoneNumber    string        `json:"phone"`
	IsAdmin        bool          `json:"is_admin"`
	Password       string        `json:"password"`
	ResetCode      string        `json:"reset_code,omitempty"`
	ResetToken     string        `json:"reset_token,omitempty"`
	CodeExpiresAt  time.Time     `json:"expires_at,omitempty"`
	HashedPassword []byte        `json:"hashed_password"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

//Exists - Returns true if a user with a phone number exists
func (user Users) Exists() bool {
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	err := collection.Find(bson.M{"phonenumber": user.PhoneNumber}).One(&user)
	if err != nil {
		return false
	}
	return true
}

//UserExists - Checks if a user with phonenumber exists
func UserExists(phoneNumber string) bool {
	var user Users
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	err := collection.Find(bson.M{"phonenumber": phoneNumber}).One(&user)
	if err != nil {
		return false
	}
	return true
}

func UserWithIDExists(id string) error {
	var user Users
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	err := collection.Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&user)
	if err != nil {
		return err
	}
	return nil
}

//Create User
func Create(user Users) error {
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	return collection.Insert(user)
}

func Read(phoneNumber string) (Users, error) {
	var user Users
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	err := collection.Find(bson.M{"phonenumber": phoneNumber}).One(&user)
	return user, err
}

//GetUserById - GetUserById
func GetUserById(id string) (Users, error) {
	var user Users
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	err := collection.Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&user)
	return user, err
}

//ReadAll - ReadAll
func ReadAll() ([]Users, error) {
	var users []Users
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	err := collection.Find(bson.M{}).All(&users)
	return users, err
}

//Update - Update
func Update(oldUser, newUser interface{}) error {
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	return collection.Update(oldUser, newUser)
}

//UpdateByID - UpdateByID
func UpdateByID(userID bson.ObjectId, user Users) error {
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.METRICSCOLLECTION)
	return collection.UpdateId(userID, user)
}

//Delete - Delete
func Delete(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid Object ID")
	}
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	return collection.Remove(bson.M{"id": bson.ObjectIdHex(id)})
}

//ConfirmResetTokens - ConfirmResetTokens
func ConfirmResetTokens(resetToken, resetCode string) (Users, error) {
	session := config.Get().Session.Clone()
	defer session.Close()
	var user Users
	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	err := collection.Find(bson.M{"resetcode": resetCode, "resettoken": resetToken}).One(&user)
	if err != nil {
		return user, errors.New("Invalid ResetCodes")
	}

	if user.CodeExpiresAt.Before(time.Now()) {
		//Time has expired
		return user, errors.New("Code has expired. Please try again")
	}
	return user, nil
}
