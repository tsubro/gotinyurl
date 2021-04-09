package main

import (
	"net/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"tiny-url/handlers"
)

const port = ":8080"

func main() {
	log.Info("Starting Tiny Url Server at PORT", port)

	httpMux := mux.NewRouter()
	httpMux.HandleFunc("/generate", handlers.GenerateTinyUrl).Methods("POST")
	httpMux.HandleFunc("/tiny/{hash}", handlers.ServeRequest)
	http.ListenAndServe(port, httpMux)
}