package problems

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/problemsApp/config"
)

type Problem struct {
	ID         bson.ObjectId `json:"_id,omitempty"`
	Title      string        `json:"title"`
	Text       string        `json:"text"`
	Pictures   []string      `json:"pictures"`
	Videos     []string      `json:"videos"`
	Document   []string      `json:"documents"`
	IsApproved bool          `json:"is_approved"` //Is the post approved by an admin or not?
	IsResolved bool          `json:"is_resolved,omitempty"`
	Status     int           `json:"status"` //0-Public 1-Private
	CreatedBy  string        `json:"created_by"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
}

//Create a Problem
func Create(problem Problem) error {
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	return collection.Insert(problem)
}

//List Problems
func ListAll() ([]Problem, error) {
	var problems []Problem
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	err := collection.Find(nil).Sort("-createdat").All(&problems)
	return problems, err
}

//ListAllApprovedListings - Lists only approved problems, meaning it can be shown to the general public if status is public
func ListAllApprovedListings() ([]Problem, error) {
	var problems []Problem
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	err := collection.Find(bson.M{"isapproved": true, "status": 0}).Sort("-createdat").All(&problems)
	return problems, err
}

//List Problems By title
func GetByTitle(title string) (Problem, error) {
	var problem Problem
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	err := collection.Find(bson.M{"title": title}).One(&problem)
	return problem, err
}

func GetByID(id string) (Problem, error) {
	var problem Problem
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	err := collection.Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&problem)
	return problem, err
}

//List Problems By a particular user
func GetUserListings(userID bson.ObjectId) ([]Problem, error) {
	session := config.Get().Session.Clone()
	defer session.Close()
	var problems []Problem

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	err := collection.Find(bson.M{"createdby": userID}).All(&problems)
	return problems, err
}

//Edit a problem by the user that created it
func Update(problemID string, update Problem) error {
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	return collection.Update(bson.M{"id": bson.ObjectIdHex(problemID)}, update)
}

func Remove(problemID string) error {
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	return collection.RemoveId(problemID)
}

func ProblemExists(id string) bool {
	session := config.Get().Session.Clone()
	defer session.Close()

	var problem Problem

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	err := collection.Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&problem)
	if err != nil {
		return false
	}
	return true
}
