package util

import (
	"github.com/prometheus/client_golang/prometheus"
)

// RequestMetricHistogram defines a type of used metrics for a specific request, using Histograms for observations
type RequestMetricHistogram struct {
	InFlightGauge prometheus.Gauge
	Counter       *prometheus.CounterVec
	Duration      *prometheus.HistogramVec
	ResponseSize  *prometheus.HistogramVec
}

// NewRequestMetricHistogram creates a RequestMetricHistogram struct with sane defaults
func NewRequestMetricHistogram(durationBuckets, responseSizeBuckets []float64) RequestMetricHistogram {
	return RequestMetricHistogram{
		InFlightGauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "in_flight_requests",
			Help: "A gauge of requests currently being served by the wrapped handler.",
		}),
		Counter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "api_requests_total",
				Help: "A counter for requests to the wrapped handler.",
			},
			[]string{"code", "method"},
		),
		Duration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "request_duration_seconds",
				Help:    "A histogram for latencies for requests.",
				Buckets: durationBuckets,
			},
			[]string{"handler", "method"},
		),
		ResponseSize: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "response_size_bytes",
				Help:    "A histogram of response sizes for requests.",
				Buckets: responseSizeBuckets,
			},
			[]string{},
		),
	}
}
