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

func runMemorySim(total uint64, usedBase uint64, cachedBase uint64, buffersBase uint64, maxDeviation float64) {
	var used, cached, buffers = usedBase, cachedBase, buffersBase

	randomStep := func(current, base uint64) uint64 {
		current += uint64((rand.Float64() - 0.5) * 60 * 1024 * 1024)
		maxDev := uint64(float64(base) * maxDeviation)
		if current < base-maxDev || current > base+maxDev {
			current = base
		}
		return current
	}

	for {
		used = randomStep(used, usedBase)
		cached = randomStep(cached, cachedBase)
		buffers = randomStep(buffers, buffersBase)

		memoryUsage.WithLabelValues("used").Set(float64(used))
		memoryUsage.WithLabelValues("cached").Set(float64(cached))
		memoryUsage.WithLabelValues("buffers").Set(float64(buffers))
		memoryUsage.WithLabelValues("free").Set(float64(total - used - cached - buffers))

		time.Sleep(100 * time.Millisecond)
	}
}
