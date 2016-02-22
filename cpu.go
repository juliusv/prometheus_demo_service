package main

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	cpuNum = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "num_cpus",
			Help:      "The number of CPUs.",
		},
	)
	cpuUsage = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "cpu_usage_seconds_total",
			Help:      "The CPU usage in seconds split by mode.",
		},
		[]string{"mode"},
	)
)

func init() {
	prometheus.MustRegister(cpuNum)
	prometheus.MustRegister(cpuUsage)
}

func runCPUSim(cpus int, userRatio, sysRatio float64) {
	cpuNum.Set(float64(cpus))
	sleep := 100 * time.Millisecond
	factor := float64(cpus) * float64(sleep) / float64(time.Second)
	for {
		userSecs := (userRatio + ((rand.Float64() - 0.5) / 5)) * factor
		sysSecs := (sysRatio + ((rand.Float64() - 0.5) / 5)) * factor
		idleSecs := factor - userSecs - sysSecs

		cpuUsage.WithLabelValues("user").Add(userSecs)
		cpuUsage.WithLabelValues("system").Add(sysSecs)
		cpuUsage.WithLabelValues("idle").Add(idleSecs)

		time.Sleep(sleep)
	}
}
