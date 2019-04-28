package likes

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/problemsApp/config"
	"github.com/ofonimefrancis/problemsApp/features/users"
)

type Likes struct {
	ID        bson.ObjectId `json:"id"`
	ProblemID bson.ObjectId `json:"problem_id,omitempty"`
	CommentID bson.ObjectId `json:"commend_id,omitempty"`
	LikedBy   users.Users   `json:"liked_by"`
	LikedOn   time.Time     `json:"liked_on"`
}

//AddLike - AddLike
func AddLike(like Likes) error {
	session := config.Get().Session.Copy()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.LIKESCOLLECTION)
	return collection.Insert(like)
}

//GetAllLikesForProblem - GetAllLikesForProblem
func GetAllLikesForProblem(problemID string) ([]Likes, error) {
	session := config.Get().Session.Copy()
	defer session.Close()
	var likes []Likes

	collection := session.DB(config.DATABASE).C(config.LIKESCOLLECTION)
	err := collection.Find(bson.M{"problemid": bson.ObjectIdHex(problemID)}).All(&likes)
	if err != nil {
		return []Likes{}, err
	}
	return likes, nil
}

//GetAllLikesForComment - GetAllLikesForComment
func GetAllLikesForComment(commentID string) ([]Likes, error) {
	session := config.Get().Session.Copy()
	defer session.Close()
	var likes []Likes

	collection := session.DB(config.DATABASE).C(config.LIKESCOLLECTION)
	err := collection.Find(bson.M{"commentid": bson.ObjectIdHex(commentID)}).All(&likes)
	if err != nil {
		return []Likes{}, err
	}
	return likes, nil
}

//DeleteLikeForComment - DeleteLikeForComment
func DeleteLikeForComment(likeID, commentID string) error {
	session := config.Get().Session.Copy()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.LIKESCOLLECTION)
	return collection.Remove(bson.M{"id": bson.ObjectIdHex(likeID), "commentid": bson.ObjectIdHex(commentID)})

}

//DeleteLikeForProblen - DeleteLikeForProblen
func DeleteLikeForProblen(likeID, problemID string) error {
	session := config.Get().Session.Copy()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.LIKESCOLLECTION)
	return collection.Remove(bson.M{"id": bson.ObjectIdHex(likeID), "problemid": bson.ObjectIdHex(problemID)})
}
