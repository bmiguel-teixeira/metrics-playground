package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	HttpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
		}, []string{"instance", "status_code", "span_name"},
	)

	HttpRequestsDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_requests_duration",
		}, []string{"instance", "status_code", "span_name"})
)

func main() {

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	for {
		generateMetrics()
		time.Sleep(10 * time.Millisecond)
	}

}

func generateMetrics() {
	clientInstances := 10
	errorRate := 10
	errorRateSkew := rand.IntN(10)

	for index := range clientInstances {
		instanceId := fmt.Sprintf("random-pod-%d", index)
		for spIndex := range clientInstances {
			spanNameId := fmt.Sprintf("/images/%d", spIndex)
			requestDurationInMs := rand.Float64() * 10
			if rand.IntN(100) < (errorRate + errorRateSkew) {
				HttpRequestsDuration.WithLabelValues(instanceId, "500", spanNameId).Observe(requestDurationInMs)
				HttpRequestsTotal.WithLabelValues(instanceId, "500", spanNameId).Inc()
			} else {
				HttpRequestsDuration.WithLabelValues(instanceId, "200", spanNameId).Observe(requestDurationInMs)
				HttpRequestsTotal.WithLabelValues(instanceId, "200", spanNameId).Inc()
			}
		}
	}
}
