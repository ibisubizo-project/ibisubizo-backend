package users

import (
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/problemsApp/config"
)

//Users - Represents a user in the system
type Users struct {
	ID             bson.ObjectId `json:"id"`
	FirstName      string        `json:"firstname"`
	MiddleName     string        `json:"middlename"`
	LastName       string        `json:"lastname"`
	PhoneNumber    string        `json:"phone"`
	IsAdmin        bool          `json:"is_admin"`
	Password       string        `json:"password"`
	HashedPassword []byte        `json:"hashed_password"`
	Salt           []byte        `json:"salt"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

//Exists - Returns true if a user with a phone number exists
func (user Users) Exists() bool {
	session := config.Get().Session
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	err := collection.Find(bson.M{"phone": user.PhoneNumber})
	if err != nil {
		return false
	}
	return true
}

func UserExists(phoneNumber string) bool {
	session := config.Get().Session
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	err := collection.Find(bson.M{"phone": phoneNumber})
	if err != nil {
		return false
	}
	return true
}

//Create User
func Create(user Users) error {
	session := config.Get().Session
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	return collection.Insert(user)
}

func Read(phoneNumber string) (Users, error) {
	var user Users
	session := config.Get().Session
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	err := collection.Find(bson.M{"phone": phoneNumber}).One(&user)
	return user, err
}

func ReadAll() ([]Users, error) {
	var users []Users
	session := config.Get().Session
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	err := collection.Find(bson.M{}).All(&users)
	return users, err
}

func Update(oldUser, newUser interface{}) error {
	session := config.Get().Session
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	return collection.Update(oldUser, newUser)
}

func Delete(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid Object ID")
	}
	session := config.Get().Session
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.USERCOLLECTION)
	return collection.Remove(bson.M{"_id": id})
}
