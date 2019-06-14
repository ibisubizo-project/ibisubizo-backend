package metrics

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type MessageResponse struct {
	Message string `json:"message"`
}

//GetProblemMetrics - GetProblemMetrics
func GetProblemMetrics(w http.ResponseWriter, r *http.Request) {
	problemID := chi.URLParam(r, "problemId")

	if !Exists(problemID) {
		log.Println("Invalid Problem ID")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: "Invalid Problem ID"})
		return
	}

	metrics, err := GetMetrics(problemID)
	if err != nil {
		log.Println(fmt.Sprintf("Unable to retrieve Metrics for problem %s", problemID))
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, MessageResponse{Message: fmt.Sprintf("Unable to retrieve Metrics for problem %s", problemID)})
		return
	}
	if len(metrics) == 0 {
		metrics = []Metrics{}
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

	metric.CreatedAt = time.Now()
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
