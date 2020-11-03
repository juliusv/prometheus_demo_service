package main

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var memoryUsage = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "memory_usage_bytes",
		Help:      "The memory usage by type.",
	},
	[]string{"type"},
)

func init() {
	prometheus.MustRegister(memoryUsage)
}

func runMemorySim(total, usedBase, cachedBase, buffersBase, maxDeviation float64) {
	var used, cached, buffers = usedBase, cachedBase, buffersBase

	randomStep := func(current, base float64) float64 {
		current += (rand.Float64() - 0.5) * 60 * 1024 * 1024
		if current < base-base*maxDeviation {
			current = base
		}
		if current > base+base*maxDeviation {
			current = base
		}
		return current
	}

	for {
		used = randomStep(used, usedBase)
		cached = randomStep(cached, cachedBase)
		buffers = randomStep(buffers, buffersBase)

		memoryUsage.WithLabelValues("used").Set(used)
		memoryUsage.WithLabelValues("cached").Set(cached)
		memoryUsage.WithLabelValues("buffers").Set(buffers)
		memoryUsage.WithLabelValues("free").Set(total - used - cached - buffers)

		time.Sleep(100 * time.Millisecond)
	}
}
