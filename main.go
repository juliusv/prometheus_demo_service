package main

import (
	"flag"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

	start = time.Now()
)

func main() {
	flag.Parse()

	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Can't parse source address", http.StatusInternalServerError)
			return
		}

		switch r.Method {
		case "GET", "POST", "PUT", "HEAD", "DELETE", "CONNECT", "OPTIONS", "NOTIFY", "TRACE", "PATCH":
			// Do nothing, allow these methods.
		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		if host != "127.0.0.1" {
			switch r.URL.Path {
			case "/api/foo", "/api/bar", "/metrics":
				// Do nothing, allow these paths.
			default:
				http.Error(w, "Not found", http.StatusNotFound)
				return
			}
		}
		handleAPI(w, r)
	})
	http.Handle("/metrics", promhttp.Handler())

	go http.ListenAndServe(*addr, nil)

	go startClient(*addr)

	go runBatchJobs(time.Minute, 10*time.Second, 0.25)
	go runCPUSim(4, 0.3, 0.2)
	go runDiskSim(160*1e9, 0.5*1e6)
	go runHolidaySim(5*time.Minute, 0.2)
	go runMemorySim(8*1024*1024*1024, 1200*1024*1024, 2500*1024*1024, 165*1024*1024, 0.5)
	prometheus.MustRegister(intermittentMetric{})

	select {}
}
