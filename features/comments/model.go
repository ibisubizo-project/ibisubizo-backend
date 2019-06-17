package comments

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/problemsApp/config"
)

type Comment struct {
	ID             bson.ObjectId `json:"_id,omitempty"`
	PostID         bson.ObjectId `json:"post_id"`
	UserID         bson.ObjectId `json:"user_id"`
	Comment        string        `json:"comment"`
	IsApproved     bool          `json:"is_approved"`
	IsAdminComment bool          `json:"is_admin"`
	Images         []string      `json:"images,omitempty"`
	CommentedAt    time.Time     `json:"commented_at"`
}

//CreateComment - CreateComment
func CreateComment(comment Comment) error {
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.COMMENTSCOLLECTION)
	return collection.Insert(comment)
}

//GetCommentsForProblem - Returns the comments for a problem listed by the latest comments
func GetCommentsForProblem(postId string) ([]Comment, error) {
	var comments []Comment
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.COMMENTSCOLLECTION)
	err := collection.Find(bson.M{"postid": bson.ObjectIdHex(postId)}).Sort("-commentedat").All(&comments)
	return comments, err
}

//GetAllUnapproved - GetAllUnapproved
func GetAllUnapproved() ([]Comment, error) {
	var comments []Comment
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.COMMENTSCOLLECTION)
	err := collection.Find(bson.M{"isapproved": false}).All(&comments)
	return comments, err
}

//GetComment - GetComment
func GetComment(commentID string) (Comment, error) {
	var comment Comment
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.COMMENTSCOLLECTION)
	err := collection.Find(bson.M{"id": bson.ObjectIdHex(commentID)}).One(&comment)
	return comment, err
}

//GetCommentsByUser - GetCommentsByUser
func GetCommentsByUser(postId, userId string) ([]Comment, error) {
	var comment []Comment
	session := config.Get().Session.Clone()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.COMMENTSCOLLECTION)
	err := collection.Find(bson.M{"postid": bson.ObjectIdHex(postId), "userid": bson.ObjectIdHex(userId)}).One(&comment)
	return comment, err
}

//CommentExists - CommentExists
func CommentExists(id string) bool {
	session := config.Get().Session.Clone()
	defer session.Close()
	var comment Comment

	collection := session.DB(config.DATABASE).C(config.COMMENTSCOLLECTION)
	err := collection.Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&comment)
	if err != nil {
		return false
	}

	return true
}

//Update - Edit a problem by the user that created it
func Update(commentID string, update Comment) error {
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.COMMENTSCOLLECTION)
	return collection.Update(bson.M{"id": bson.ObjectIdHex(commentID)}, update)
}

//Remove - Remove
func Remove(problemID string) error {
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.COMMENTSCOLLECTION)
	return collection.Remove(bson.M{"id": bson.ObjectIdHex(problemID)})
}
