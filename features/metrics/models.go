package metrics

import (
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/problemsApp/config"
)

//Metrics - Keeps track of Site visits
type Metrics struct {
	ProblemID string    `json:"problem_id"`
	CreatedAt time.Time `json:"created_at"`
	Visits    int       `json:"visits"`
	Year      int       `json:"year"`
	Month     int       `json:"month"`
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
func GetMetrics(problemID string, month, year int) (Metrics, error) {
	session := config.Get().Session.Clone()
	defer session.Close()
	var metrics Metrics
	collection := session.DB(config.DATABASE).C(config.METRICSCOLLECTION)
	err := collection.Find(bson.M{"problemid": problemID, "month": month, "year": year}).One(&metrics)
	return metrics, err
}

//GetAllMetrics - GetAllMetrics
func GetAllMetrics() ([]Metrics, error) {
	session := config.Get().Session.Clone()
	defer session.Close()
	var metrics []Metrics

	collection := session.DB(config.DATABASE).C(config.METRICSCOLLECTION)
	err := collection.Find(bson.M{}).All(&metrics)
	return metrics, err
}

//Update - Update
func Update(problemID string, metrics Metrics) error {
	session := config.Get().Session.Clone()
	defer session.Close()
	collection := session.DB(config.DATABASE).C(config.METRICSCOLLECTION)
	return collection.Update(bson.M{"problemid": problemID, "year": metrics.Year, "month": metrics.Month}, metrics)
}

//GetMonthlyMetrics - GetMonthlyMetrics
func GetMonthlyMetrics(month, year int) ([]Metrics, error) {
	session := config.Get().Session.Clone()
	defer session.Close()

	var metrics []Metrics
	collection := session.DB(config.DATABASE).C(config.METRICSCOLLECTION)
	err := collection.Find(bson.M{"month": month, "year": year}).All(&metrics)
	if err != nil {
		log.Println(err)
		return []Metrics{}, err
	}
	return metrics, nil
}
