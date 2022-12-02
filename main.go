package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Customer struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customers = map[string]Customer{
	"aaa": {
		Name:      "Aye",
		Role:      "AyeAye",
		Email:     "Aye@local.host",
		Phone:     "(123) 456-7893",
		Contacted: true,
	},
	"aab": {
		Name:      "Bay",
		Role:      "BayBay",
		Email:     "Bay@local.host",
		Phone:     "(123)456-7892",
		Contacted: false,
	},
	"aac": {
		Name:      "Cey",
		Role:      "CeyCey",
		Email:     "Cey@local.host",
		Phone:     "(123) 456-7891",
		Contacted: false,
	},
	"aad": {
		Name:      "Dey",
		Role:      "DeyDey",
		Email:     "Dey@local.host",
		Phone:     "(123) 456-7899",
		Contacted: false,
	},
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if _, ok := customers[mux.Vars(r)["id"]]; ok {
		json.NewEncoder(w).Encode(customers[mux.Vars(r)["id"]])
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func addCustomer(w http.ResponseWriter, r *http.Request) {

	var customer Customer
	customer.Id = uuid.New().String()
	reqBody, _ := ioutil.ReadAll(r.Body)

	w.Header().Set("Content-Type", "application/json")

	if _, ok := customers[mux.Vars(r)["id"]]; ok {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(customers)
	} else {

		json.Unmarshal(reqBody, &customer)
		customers[customer.Id] = customer
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(customers)
	}

}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	// delete an existing customer
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	if _, ok := customers[id]; ok {
		delete(customers, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	// update an existing customer
	w.Header().Set("Content-Type", "application/json")
	var customer Customer
	reqBody, _ := ioutil.ReadAll(r.Body)
	id := mux.Vars(r)["id"]
	if _, ok := customers[id]; ok {
		json.Unmarshal(reqBody, &customer)
		customers[id] = customer
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}
}
func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()

	if port == "" {
		log.Printf("$PORT variable not set. Setting to default 8080")
		port = "8080"
	}

	log.Printf("Starting the Server on %s...", port)
	router.HandleFunc("/customers", getCustomers).Methods("GET") // get all customers
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT") // update a customer

	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE") // delete a customer

	log.Fatal(http.ListenAndServe(":"+port, router))

}
