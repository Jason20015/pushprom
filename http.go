package main

import (
	//"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func listenHTTP() {
	log.Printf("exposing metrics on http://" + *httpListenAddress + "/metrics\n")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*httpListenAddress, nil))
}
