package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HealthCheckHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}
