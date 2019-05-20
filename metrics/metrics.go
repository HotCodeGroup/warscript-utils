package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// NewHits получение новой метрики по хитам
func NewHits() *prometheus.CounterVec {
	hits := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"status", "path"})
	prometheus.MustRegister(hits)

	return hits
}
