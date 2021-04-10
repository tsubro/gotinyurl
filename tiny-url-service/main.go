package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"tiny-url/handlers"
	log "github.com/sirupsen/logrus"
)

const port = ":8080"

func main() {
	log.Info("Starting Tiny Url Server at PORT", port)

	httpMux := mux.NewRouter()
	httpMux.HandleFunc("/encode", handlers.EncodeUrl).Methods("POST")
	httpMux.HandleFunc("/{hash}/info", handlers.Info).Methods("GET")
	httpMux.HandleFunc("/{hash}", handlers.ServeRequest).Methods("GET")
	http.ListenAndServe(port, httpMux)
}