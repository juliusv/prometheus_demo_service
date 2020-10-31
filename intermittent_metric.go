package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type intermittentMetric struct{}

var intermittentMetricDesc = prometheus.NewDesc("demo_intermittent_metric", "A metric that is only present intermittently to test staleness handling.", []string{}, nil)

func (e intermittentMetric) Describe(ch chan<- *prometheus.Desc) {
	ch <- intermittentMetricDesc
}

func (e intermittentMetric) Collect(ch chan<- prometheus.Metric) {
	// Expose this metric only every second minute.
	if time.Now().Minute()%2 == 0 {
		ch <- prometheus.MustNewConstMetric(intermittentMetricDesc, prometheus.GaugeValue, 1)
	}
}
