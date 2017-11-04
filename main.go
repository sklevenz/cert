package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"fmt"
)

type Domain struct {
	Name     string   `json:"name"`
	Subject  string   `json: "subject"`
	AltNames []string `json:"altnames"`
}

var domains []Domain

// our main function
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/domains", GetDomains).Methods("GET")
	router.HandleFunc("/domains/{id}", GetDomain).Methods("GET")
	router.HandleFunc("/domains/{id}", CreateDomain).Methods("POST")
	router.HandleFunc("/domains/{id}", DeleteDomain).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("PORT:", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func GetDomains(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(domains)
}

func GetDomain(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range domains {
		if item.Name == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Domain{})
}

func CreateDomain(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var domain Domain
	_ = json.NewDecoder(r.Body).Decode(&domain)
	domain.Name = params["id"]
	domains = append(domains, domain)
	json.NewEncoder(w).Encode(domains)
}

func DeleteDomain(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range domains {
		if item.Name == params["id"] {
			domains = append(domains[:index], domains[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(domains)
	}
}
