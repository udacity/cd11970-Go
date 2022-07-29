package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var members = map[string]string{
	"1": "Andy",
	"2": "Peter",
	"3": "Gabriella",
	"4": "Jordy",
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(members)
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, ok := members[id]; ok {
		delete(members, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(members)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(members)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/members", getMembers).Methods("GET")
	router.HandleFunc("/deleteMember/{id}", deleteMember).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
