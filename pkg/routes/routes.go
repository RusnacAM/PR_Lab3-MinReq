package routes

import (
	"Lab3/pkg/Datastore"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/datastore", getAll).Methods("GET")
	router.HandleFunc("/datastore/{id}", getByID).Methods("GET")
	router.HandleFunc("/datastore", create).Methods("POST")
	router.HandleFunc("/datastore/{id}", update).Methods("PUT")
	router.HandleFunc("/datastore/{id}", deleteByID).Methods("DELETE")
}

func getAll(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(Datastore.DataStore)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func getByID(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var ID = vars["id"]

	value, err := Datastore.Read(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else {
		res, _ := json.Marshal(value)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	value, err := Datastore.Create(string(body))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	} else {
		res, _ := json.Marshal(value)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var ID = vars["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	value, err := Datastore.Update(ID, string(body))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else {
		res, _ := json.Marshal(value)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func deleteByID(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var ID = vars["id"]

	value, err := Datastore.Delete(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else {
		res, _ := json.Marshal(value)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
