package metrics

import "github.com/go-chi/chi"

const (
	//ProblemMetricsRoute - ProblemMetricsRoute
	ProblemMetricsRoute = "/{problemId}"
	//AddMetricsRoute - AddMetricsRoute
	AddMetricsRoute = "/"
)

//Routes -  All the metrics specific routes
func Routes() *chi.Mux {
	router := chi.NewMux()
	router.Get(ProblemMetricsRoute, GetProblemMetrics)
	router.Post(AddMetricsRoute, NewMetrics)
	router.Put("/{problemId}", UpdateMetrics)
	return router
}
