package metrics

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ofonimefrancis/problemsApp/features/problems"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

//MessageResponse - MessageResponse
type MessageResponse struct {
	Message string `json:"message"`
}

//GetProblemMetrics - GetProblemMetrics
func GetProblemMetrics(w http.ResponseWriter, r *http.Request) {
	problemID := chi.URLParam(r, "problemId")
	currentYear := time.Now().Year()
	currentMonth := int(time.Now().Month())

	if !Exists(problemID) {
		log.Println("Invalid Problem ID")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Problem ID"})
		return
	}

	metrics, err := GetMetrics(problemID, currentMonth, currentYear)
	if err != nil {
		log.Println(fmt.Sprintf("Unable to retrieve Metrics for problem %s", problemID))
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: fmt.Sprintf("Unable to retrieve Metrics for problem %s", problemID)})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, metrics)
}

//NewMetrics - NewMetrics
func NewMetrics(w http.ResponseWriter, r *http.Request) {
	var metric Metrics
	err := json.NewDecoder(r.Body).Decode(&metric)
	if err != nil {
		log.Println("Invalid Payload")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Payload"})
		return
	}
	currentMonthString := time.Now().Month()
	if err != nil {
		log.Println(err)
	}
	metric.CreatedAt = time.Now()
	metric.Year = time.Now().Year()
	metric.Month = int(currentMonthString)

	log.Println("Current month string ", currentMonthString)
	log.Println("Metrics to add ")
	log.Println(metric)

	err = AddMetrics(metric)
	if err != nil {
		log.Println(err)
		log.Println("Error adding a new metrics")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error adding a new metrics"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, MessageResponse{Message: "Metrics successfully added"})
	return

}

//UpdateMetrics - UpdateMetrics
func UpdateMetrics(w http.ResponseWriter, r *http.Request) {
	problemID := chi.URLParam(r, "problemId")
	if !Exists(problemID) {
		log.Println("Problem metrics not found...")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Metric For Problem Not Found"})
		return
	}
	var metric Metrics

	err := json.NewDecoder(r.Body).Decode(&metric)
	if err != nil {
		log.Println(err)
		log.Println("Invalid Payload")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Payload"})
		return
	}

	err = Update(problemID, metric)
	if err != nil {
		log.Println(err)
		log.Println("Error Updating Problem Metrics")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error Updating Problem Metrics"})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, MessageResponse{Message: "Metrics Updated"})
	return
}

//TrendingProblems - TrendingProblems
func TrendingProblems(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	currentMonth := int(now.Month())
	metrics, err := GetMonthlyMetrics(currentMonth, now.Year())
	if err != nil {
		log.Println("Error retrieving trending problems")
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Error retrieving trending problems"})
		return
	}

	trendingProblems := []problems.Problem{}

	for _, metric := range metrics {
		problem, err := problems.GetByID(metric.ProblemID)
		if err != nil {
			continue
		}
		trendingProblems = append(trendingProblems, problem)
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, trendingProblems)
	return
}
