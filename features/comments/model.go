package comments

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/problemsApp/config"
)

type Comment struct {
	ID          bson.ObjectId `json:"_id,omitempty"`
	PostID      bson.ObjectId `json:"post_id"`
	UserID      bson.ObjectId `json:"user_id"`
	Comment     string        `json:"comment"`
	Images      []string      `json:"images,omitempty"`
	CommentedAt time.Time     `json:"commented_at"`
}

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
