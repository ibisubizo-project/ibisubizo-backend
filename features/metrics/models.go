package metrics

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/problemsApp/config"
)

//Metrics - Keeps track of Site visits
type Metrics struct {
	ProblemID string    `json:"problem_id"`
	Visits    int       `json:"visits"`
	CreatedAt time.Time `json:"created_at"`
}

//Exists - Checks if a problem has metrics
func Exists(problemID string) bool {
	session := config.Get().Session.Clone()
	defer session.Close()
	var metrics Metrics

	collection := session.DB(config.DATABASE).C(config.METRICSCOLLECTION)
	err := collection.Find(bson.M{"problemid": problemID}).One(&metrics)
	if err != nil {
		return false
	}
	return true
}

//AddMetrics - Add a Metrics
func AddMetrics(metric Metrics) error {
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.METRICSCOLLECTION)
	return collection.Insert(metric)
}

//GetMetrics - GetMetrics
func GetMetrics(problemID string) ([]Metrics, error) {
	session := config.Get().Session.Clone()
	defer session.Close()
	var metrics []Metrics
	collection := session.DB(config.DATABASE).C(config.METRICSCOLLECTION)
	err := collection.Find(bson.M{"problemid": problemID}).All(&metrics)
	return metrics, err
}
