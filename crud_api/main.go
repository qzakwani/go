package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title     string
	CreatedOn time.Time
}

var notes = make(map[string]Note)
var id int

func postNote(w http.ResponseWriter, r *http.Request) {
	id++
	var note Note
	json.NewDecoder(r.Body).Decode(&note)
	note.CreatedOn = time.Now()
	notes[strconv.Itoa(id)] = note

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	res, _ := json.Marshal(&note)
	// json.NewEncoder(w).Encode(res)
	w.Write(res)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/", postNote).Methods("POST")
	r.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(notes[mux.Vars(r)["id"]])
		w.Write(res)
	}).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	server.ListenAndServe()
}
