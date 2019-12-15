package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var x map[string]int

type Order struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var orders []Order

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("ok")
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	fmt.Println("CreateOrder", order, r.Body)
	_ = json.NewDecoder(r.Body).Decode(&order)
	order.ID = strconv.Itoa(rand.Intn(1000000))
	orders = append(orders, order)
	json.NewEncoder(w).Encode(&order)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getOrder")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range orders {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&Order{})
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println("delete Order", orders)
	for index, item := range orders {
		fmt.Println("index", index)
		fmt.Println("item", item.ID)
		fmt.Println("params", params["id"])
		fmt.Println("params", orders[index])
		if item.ID == params["id"] {
			// orders = append(orders[:index], orders[index+1]...)
			break
		}
	}
	json.NewEncoder(w).Encode(orders)
}

func main() {
	fmt.Println("It works!")
	router := mux.NewRouter()
	router.HandleFunc("/", health).Methods("GET")
	router.HandleFunc("/order", createOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", getOrder).Methods("GET")
	router.HandleFunc("/orders/{id}", deleteOrder).Methods("DELETE")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	// http.ListenAndServe(":8000", router)
}
