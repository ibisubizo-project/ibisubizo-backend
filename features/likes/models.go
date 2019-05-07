package likes

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/problemsApp/config"
)

type ProblemLikes struct {
	ID        bson.ObjectId `json:"id"`
	ProblemID bson.ObjectId `json:"problem_id,omitempty"`
	//CommentID bson.ObjectId `json:"commend_id,omitempty"`
	LikedBy bson.ObjectId `json:"liked_by"`
	LikedOn time.Time     `json:"liked_on"`
}

type CommentLikes struct {
	ID bson.ObjectId `json:"id"`
	//ProblemID bson.ObjectId `json:"problem_id,omitempty"`
	CommentID bson.ObjectId `json:"comment_id,omitempty"`
	LikedBy   bson.ObjectId `json:"liked_by"`
	LikedOn   time.Time     `json:"liked_on"`
}

//AddLike - AddLike
func AddProblemLike(like ProblemLikes) error {
	session := config.Get().Session.Copy()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMLIKESCOLLECTION)
	return collection.Insert(like)
}

func AddCommentLike(like CommentLikes) error {
	session := config.Get().Session.Copy()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.COMMENTLIKESCOLLECTION)
	return collection.Insert(like)
}

//GetAllLikesForProblem - GetAllLikesForProblem
func GetAllLikesForProblem(problemID string) ([]ProblemLikes, error) {
	session := config.Get().Session.Copy()
	defer session.Close()
	var likes []ProblemLikes

	collection := session.DB(config.DATABASE).C(config.PROBLEMLIKESCOLLECTION)
	err := collection.Find(bson.M{"problemid": bson.ObjectIdHex(problemID)}).All(&likes)
	if err != nil {
		return []ProblemLikes{}, err
	}
	return likes, nil
}

//GetAllLikesForComment - GetAllLikesForComment
func GetAllLikesForComment(commentID string) ([]CommentLikes, error) {
	session := config.Get().Session.Copy()
	defer session.Close()
	var likes []CommentLikes

	collection := session.DB(config.DATABASE).C(config.COMMENTLIKESCOLLECTION)
	err := collection.Find(bson.M{"commentid": bson.ObjectIdHex(commentID)}).All(&likes)
	if err != nil {
		return []CommentLikes{}, err
	}
	return likes, nil
}

//DeleteLikeForComment - DeleteLikeForComment
func DeleteLikeForComment(likeID, commentID string) error {
	session := config.Get().Session.Copy()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.COMMENTLIKESCOLLECTION)
	return collection.Remove(bson.M{"id": bson.ObjectIdHex(likeID), "commentid": bson.ObjectIdHex(commentID)})

}

//DeleteLikeForProblen - DeleteLikeForProblen
func DeleteLikeForProblem(problemID, likedBy string) error {
	session := config.Get().Session.Copy()
	defer session.Close()

	collection := session.DB(config.DATABASE).C(config.PROBLEMLIKESCOLLECTION)
	return collection.Remove(bson.M{"likedby": bson.ObjectIdHex(likedBy), "problemid": bson.ObjectIdHex(problemID)})
}
