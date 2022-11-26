package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getAll(w http.ResponseWriter, r *http.Request) {

}

func getByID(w http.ResponseWriter, r *http.Request) {

}

func create(w http.ResponseWriter, r *http.Request) {

}

func update(w http.ResponseWriter, r *http.Request) {

}

func delete(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/datastore", getAll).Methods("GET")
	router.HandleFunc("/datastore/{id}", getByID).Methods("GET")
	router.HandleFunc("/datastore", create).Methods("POST")
	router.HandleFunc("/datastore/{id}", update).Methods("PUT")
	router.HandleFunc("/datastore/{id}", delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
