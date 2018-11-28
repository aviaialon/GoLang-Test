package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"request"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/multisite/create", MultiSiteCreate).Methods("GET")
	router.HandleFunc("/multisite/update", MultiSiteUpdate).Methods("GET")
	log.Print("Started server on port 8000");
	log.Fatal(http.ListenAndServe(":8000", router))
}

func MultiSiteCreate(w http.ResponseWriter, r *http.Request) {
	respond(w, request.NewFromRequest(r))
}

func MultiSiteUpdate(w http.ResponseWriter, r *http.Request) {}

func respond(w http.ResponseWriter, data interface{})  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(request.All())
}