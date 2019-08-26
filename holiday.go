package main

import (
	"math"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	isHoliday = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "is_holiday",
			Help:      "Set to 1 if it is currently a holiday, 0 otherwise.",
		})

	shippedItems = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "items_shipped_total",
			Help:      "The total number of shipped items. Affected heavily by whether it's currently a holiday.",
		})

	isHolidayVar = uint64(0)
)

func init() {
	prometheus.MustRegister(isHoliday)
	prometheus.MustRegister(shippedItems)
}

func runHolidaySim(dayLength time.Duration, holidayRatio float64) {
	start := time.Now()
	go func() {
		for {
			shippedItems.Inc()
			factor := 2 + math.Sin(1+2*math.Pi*float64(time.Since(start))/float64(dayLength))
			d := 100 * factor
			if isHolidayVar == 1 {
				d *= 1.8
			}
			time.Sleep(time.Duration(d) * time.Millisecond)
		}
	}()

	ticker := time.NewTicker(dayLength)
	for {
		if rand.Float64() > holidayRatio {
			isHoliday.Set(0)
			atomic.StoreUint64(&isHolidayVar, 0)
		} else {
			isHoliday.Set(1)
			atomic.StoreUint64(&isHolidayVar, 1)
		}

		<-ticker.C
	}
}
