package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var people []Person

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/person", PersonHandler).Methods("POST")
	r.HandleFunc("/list", ListHandler).Methods("GET")
	r.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)
	productJSON, err := json.Marshal(people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(productJSON)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(people) == 0 {
		http.Error(w, "No person found", http.StatusNotFound)
		return
	}
	productJSON, err := json.Marshal(people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(productJSON)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("params: ",params)
	for index, item := range people {
		log.Println("item",item)
		log.Println("item.ID",item.ID)
		if item.ID == params["id"] {
			log.Println("Removing person with id: ", item.ID)
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
