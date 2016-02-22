package main

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	diskUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "disk",
			Name:      "usage_bytes",
			Help:      "Disk usage in bytes.",
		},
	)
	diskTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "disk",
			Name:      "total_bytes",
			Help:      "Total disk space in bytes.",
		},
	)
)

func init() {
	prometheus.MustRegister(diskUsage)
	prometheus.MustRegister(diskTotal)
}

func runDiskSim(total int64, inc int64) {
	diskTotal.Set(float64(total))
	usage := int64(0.1 * float64(total))
	for {
		usage += int64(float64(inc) * (1 - (rand.Float64() - 0.5)))

		if float64(usage) > 0.9*float64(total) {
			usage = int64(0.1 * float64(total))
		}
		diskUsage.Set(float64(usage))

		time.Sleep(time.Second)
	}
}
