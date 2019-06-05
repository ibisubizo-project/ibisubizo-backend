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
func ListAll(page int) ([]Problem, error) {
	var problems []Problem
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	var err error

	if page != 0 {
		err = collection.Find(nil).Sort("-createdat").Skip((page - 1) * config.LIMITS).Limit(config.LIMITS).All(&problems)
	} else {
		err = collection.Find(nil).Sort("-createdat").All(&problems)
	}

	return problems, err
}

//ListAllApprovedListings - Lists only approved problems, meaning it can be shown to the general public if status is public
func ListAllApprovedListings(page int) ([]Problem, error) {
	var problems []Problem
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	var err error

	if page != 0 {
		err = collection.Find(bson.M{"isapproved": true, "status": 0}).Sort("-createdat").Skip((page - 1) * config.LIMITS).Limit(config.LIMITS).All(&problems)
	} else {
		err = collection.Find(bson.M{"isapproved": true, "status": 0}).Sort("-createdat").All(&problems)
	}

	return problems, err
}

//ListAllResolvedProblems - ListAllResolvedProblems
func ListAllResolvedProblems(page int) ([]Problem, error) {
	var problems []Problem
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	var err error

	if page != 0 {
		err = collection.Find(bson.M{"isresolved": true}).Sort("-createdat").Skip((page - 1) * config.LIMITS).Limit(config.LIMITS).All(&problems)
	} else {
		err = collection.Find(bson.M{"isresolved": true}).Sort("-createdat").All(&problems)
	}

	return problems, err
}

//ListAllUnResolvedProblems - ListAllUnResolvedProblems
func ListAllUnResolvedProblems(page int) ([]Problem, error) {
	var problems []Problem
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	var err error

	if page != 0 {
		err = collection.Find(bson.M{"isresolved": false}).Sort("-createdat").Skip((page - 1) * config.LIMITS).Limit(config.LIMITS).All(&problems)
	} else {
		err = collection.Find(bson.M{"isresolved": false}).Sort("-createdat").All(&problems)
	}

	return problems, err
}

//ListAllUnApprovedProblems - ListAllUnApprovedProblems
func ListAllUnApprovedProblems(page int) ([]Problem, error) {
	var problems []Problem
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	var err error

	if page != 0 {
		err = collection.Find(bson.M{"status": 0, "isapproved": false}).Sort("-createdat").Skip((page - 1) * config.LIMITS).Limit(config.LIMITS).All(&problems)
	} else {
		err = collection.Find(bson.M{"status": 0, "isapproved": false}).Sort("-createdat").All(&problems)
	}

	return problems, err
}

//GetByTitle - List Problems By title
func GetByTitle(title string) (Problem, error) {
	var problem Problem
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	err := collection.Find(bson.M{"title": title}).One(&problem)
	return problem, err
}

//GetByID - GetByID
func GetByID(id string) (Problem, error) {
	var problem Problem
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	err := collection.Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&problem)
	return problem, err
}

//GetUsersProblemByID - GetUsersProblemByID
func GetUsersProblemByID(problemID, userID string) (Problem, error) {
	var problem Problem
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	err := collection.Find(bson.M{"id": bson.ObjectIdHex(problemID), "createdby": bson.IsObjectIdHex(userID)}).One(&problem)
	return problem, err
}

//GetUserListings - List Problems By a particular user
func GetUserListings(userID string, page int) ([]Problem, error) {
	session := config.Get().Session.Clone()
	defer session.Close()
	var problems []Problem
	var err error

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	if page != 0 {
		err = collection.Find(bson.M{"createdby": userID}).Sort("-createdat").Skip((page - 1) * config.LIMITS).Limit(config.LIMITS).All(&problems)
	} else {
		err = collection.Find(bson.M{"createdby": userID}).Sort("-createdat").All(&problems)
	}
	return problems, err
}

//Edit a problem by the user that created it
func Update(problemID string, update Problem) error {
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	return collection.Update(bson.M{"id": bson.ObjectIdHex(problemID)}, update)
}

//Remove - Remove
func Remove(problemID string) error {
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.PROBLEMSCOLLECTION)
	return collection.Remove(bson.M{"id": bson.ObjectIdHex(problemID)})
}

//ProblemExists - ProblemExists
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
