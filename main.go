package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

	start = time.Now()
)

func main() {
	flag.Parse()

	http.HandleFunc("/api/", handleAPI)
	http.Handle("/metrics", prometheus.Handler())

	go http.ListenAndServe(*addr, nil)

	go startClient(*addr)

	go runBatchJobs(time.Minute, 10*time.Second, 0.25)
	go runCPUSim(4, 0.3, 0.2)
	go runDiskSim(160*1e9, 0.5*1e6)

	select {}
}
